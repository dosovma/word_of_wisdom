package tcp

import (
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"

	"server/internal/service"
	"server/internal/service/entity"
	"server/pkg/logger"
	"server/pkg/tcp"
)

type tokenStorage interface {
	Token(tokenID uuid.UUID) (*entity.Token, error)
	Store(entity.Token)
}

type messenger interface {
	Write(messages []string) error
	Read() ([]string, error)
}

type Handler struct {
	service service.IService
	m       messenger
	auth    tokenStorage
	log     logger.Logger
}

func NewHandler(service service.IService, storage tokenStorage, log logger.Logger) *Handler {
	return &Handler{
		service: service,
		auth:    storage,
		log:     log,
	}
}

func (h *Handler) Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			h.log.Printf("failed to close connection: %s\n", err)
		}
		h.log.Println("connection closed")
	}(conn)

	h.m = tcp.NewMessenger(conn, MESSAGE_START, MESSAGE_END, MESSAGE_SIZE_LIMIT)

	for {
		h.log.Println("start handling")

		msg, err := h.m.Read()
		if err != nil {
			h.log.Println("failed to read message")
			return // write error
		}

		h.log.Println("message read")

		cmd, err := tcp.GetDataByHeader(COMMAND, msg)
		if err != nil {
			h.log.Println("failed to read command")
			return
		}

		switch cmd {
		case CMD_TOKEN:
			h.log.Println("got token request")

			request, err := extractRequest(msg)
			if err != nil {
				h.log.Println("failed to get request header")
				return // invalid response format
			}

			challenge := h.service.Challenge(*request)

			if err = h.m.Write([]string{CHALLENGE + challenge}); err != nil {
				h.log.Println("failed to send challenge")
				return //
			}
			h.log.Println("challenge sent")
		case CMD_SOLUTION:
			h.log.Println("got validation request")

			solution, err := tcp.GetDataByHeader(SOLUTION, msg)
			if err != nil {
				h.log.Println("failed to get solution header")
				return
			}

			if isGranted := h.service.Validate(solution); !isGranted {
				if err = h.m.Write([]string{ACCESS + "Reject"}); err != nil {
					return // TODO
				}

				continue
			}

			token := h.service.Token()

			if err = h.m.Write([]string{TOKEN + token.String()}); err != nil {
				h.log.Println("failed to send token")
				return
			}
		case CMD_QUOTE:
			h.log.Println("got quote request")

			if h.Auth(msg) {
				q, err := h.service.Quote()
				if err != nil {
					h.log.Println("failed to get quote from service")
					return //
				}

				if err = h.m.Write([]string{QUOTE + q}); err != nil {
					h.log.Println("failed to send quote")
					return
				}

				continue
			}

			if err = h.m.Write([]string{ACCESS + "Reject"}); err != nil {
				return // TODO
			}
		}
	}
}

func extractRequest(request []string) (*entity.Request, error) {
	idStr, err := tcp.GetDataByHeader(REQUEST_ID, request)
	if err != nil {
		return nil, err // invalid response format
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err // invalid response format
	}

	timeStr, err := tcp.GetDataByHeader(REQUEST_TIME, request)
	if err != nil {
		return nil, err // invalid response format
	}

	t, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return nil, err // invalid response format
	}

	return &entity.Request{
		ID:        id,
		CreatedAt: t,
	}, nil
}

func (h *Handler) Auth(messages []string) bool {
	tokenStr, err := tcp.GetDataByHeader(TOKEN, messages)
	if err != nil {
		return false
	}

	tokenID, err := uuid.Parse(tokenStr)
	if err != nil {
		return false
	}

	token, err := h.auth.Token(tokenID)
	if err != nil {
		h.log.Printf("failed to get token: %s", err)

		return false
	}

	return token.ExpiryDate.After(time.Now())
}
