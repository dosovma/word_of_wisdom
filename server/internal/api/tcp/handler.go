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
	Write(conn net.Conn, messages []string) error
	Read(conn net.Conn) ([]string, error)
}

type Handler struct {
	service service.IService
	m       messenger
	auth    tokenStorage
	log     logger.Logger
}

func NewHandler(service service.IService, storage tokenStorage, m messenger, log logger.Logger) *Handler {
	return &Handler{
		service: service,
		auth:    storage,
		log:     log,
		m:       m,
	}
}

func (h *Handler) Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			h.log.Printf("failed to close connection: %s\n", err)
		}
		h.log.Println("connection closed")
	}(conn)

	for {
		h.log.Println("start handling")

		msg, err := h.m.Read(conn)
		if err != nil {
			h.log.Println("failed to read message")
			return // write error
		}

		h.log.Println("message read")

		cmd, err := tcp.GetDataByHeader(Command, msg)
		if err != nil {
			h.log.Println("failed to read command")
			return
		}

		switch cmd {
		case CmdToken:
			h.log.Println("got token request")

			request, err := extractRequest(msg)
			if err != nil {
				h.log.Println("failed to get request header")

				return
			}

			challenge := h.service.Challenge(*request)

			if err = h.m.Write(conn, []string{Challenge + challenge}); err != nil {
				h.log.Println("failed to send challenge")
				return //
			}
			h.log.Println("challenge sent")
		case CmdSolution:
			h.log.Println("got validation request")

			solution, err := tcp.GetDataByHeader(Solution, msg)
			if err != nil {
				h.log.Println("failed to get solution header")
				return
			}

			if isGranted := h.service.Validate(solution); !isGranted {
				if err = h.m.Write(conn, []string{Access + "Reject"}); err != nil {
					return
				}

				continue
			}

			token := h.service.Token()

			if err = h.m.Write(conn, []string{Token + token.String()}); err != nil {
				h.log.Println("failed to send token")
				return
			}
		case CmdQuote:
			h.log.Println("got quote request")

			if h.Auth(msg) {
				q, err := h.service.Quote()
				if err != nil {
					h.log.Println("failed to get quote from service")
					return //
				}

				if err = h.m.Write(conn, []string{Quote + q}); err != nil {
					h.log.Println("failed to send quote")

					return
				}

				return // close connection
			}

			if err = h.m.Write(conn, []string{Access + "Reject"}); err != nil {
				return
			}
		}
	}
}

func extractRequest(request []string) (*entity.Request, error) {
	idStr, err := tcp.GetDataByHeader(RequestID, request)
	if err != nil {
		return nil, err // invalid response format
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err // invalid response format
	}

	timeStr, err := tcp.GetDataByHeader(RequestTime, request)
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
	tokenStr, err := tcp.GetDataByHeader(Token, messages)
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
