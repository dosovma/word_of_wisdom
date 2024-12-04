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
	Write(header string, messages []string) error
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
	}(conn)

	h.m = tcp.NewMessenger(conn, MESSAGE_START, MESSAGE_END, MESSAGE_SIZE_LIMIT)

	for {
		msg, err := h.m.Read()
		if err != nil {
			return // write error
		}

		cmd, err := tcp.GetDataByHeader(COMMAND, msg)
		if err != nil {
			return
		}

		switch cmd {
		case CMD_TOKEN:
			request, err := extractRequest(msg)
			if err != nil {
				return // invalid response format
			}

			challenge := h.service.Challenge(*request)

			if err = h.m.Write(CHALLENGE, []string{challenge}); err != nil {
				return //
			}
		case CMD_SOLUTION:
			solution, err := tcp.GetDataByHeader(SOLUTION, msg)
			if err != nil {

			}

			if isGranted := h.service.Validate(solution); !isGranted {
				if err = h.m.Write(ACCESS, []string{"Reject"}); err != nil {
					return // TODO
				}

				continue
			}

			token := h.service.Token()

			if err = h.m.Write(TOKEN, []string{token.String()}); err != nil {
				return // TODO
			}
		case CMD_QUOTE:
			if h.Auth(msg) {
				q, err := h.service.Quote()
				if err != nil {
					return //
				}

				if err = h.m.Write(QUOTE, []string{q}); err != nil {
					return //
				}

				continue
			}

			if err = h.m.Write(ACCESS, []string{"Reject"}); err != nil {
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

	return token.ExpiryDate.Before(time.Now())
}
