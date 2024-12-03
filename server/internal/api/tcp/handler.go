package tcp

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"
	"server/internal/service"
	"server/internal/service/entity"
	"server/pkg/tcp"
)

type tokenStorage interface {
	Token(tokenID uuid.UUID) entity.Token
	Store(entity.Token)
}

type Handler struct {
	service service.IService
	auth    tokenStorage
}

func NewHandler(service service.IService, storage tokenStorage) *Handler {
	return &Handler{
		service: service,
		auth:    storage,
	}
}

func (h *Handler) Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			// log
		}
	}(conn)

	for {
		msg, err := messageReader(conn)
		if err != nil {
			return // write error
		}

		cmd, err := tcp.GetDataByHeader(COMMAND, msg)
		if err != nil {
			return
		}

		switch cmd {
		case CMD_REQUEST:
			if h.Auth(msg) {
				q, err := h.service.Quote()
				if err != nil {
					return //
				}

				if err = writeMessage(conn, QUOTE, []string{q}); err != nil {
					return //
				}

				continue
			}

			request, err := extractRequest(msg)
			if err != nil {
				return // invalid response format
			}

			challenge := h.service.Challenge(*request)

			if err = writeMessage(conn, CHALLENGE, []string{challenge}); err != nil {
				return //
			}
		case CMD_SOLUTION:
			solution, err := tcp.GetDataByHeader(SOLUTION, msg)
			if err != nil {

			}

			if isGranted := h.service.Validate(solution); !isGranted {
				if err = writeMessage(conn, ACCESS, []string{"Reject"}); err != nil {
					return // TODO
				}

				continue
			}

			token := h.service.Token()

			if err = writeMessage(conn, TOKEN, []string{token.String()}); err != nil {
				return // TODO
			}
		}
	}
}

func writeMessage(conn net.Conn, header string, messages []string) error {
	message := make([]string, 0, len(messages)+2)
	message = append(message, MESSAGE_START)
	message = append(message, tcp.AddDataToHeader(header, messages)...)
	message = append(message, MESSAGE_END)

	for _, msg := range message {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			return err // TODO
		}
	}

	return nil
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

func messageReader(conn net.Conn) ([]string, error) {
	messageSize := 0
	connReader := bufio.NewReader(conn)

	message := make([]string, 0, 1)

	for {
		str, err := connReader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		messageSize += len(str)
		if messageSize >= MESSAGE_SIZE_LIMIT {
			return nil, errors.New("message limit exceeded")
		}

		if isPayload(str) {
			message = append(message, str)
		}

		if str == MESSAGE_END {
			return message, nil
		}
	}
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

	token := h.auth.Token(tokenID)

	return token.ExpiryDate.Before(time.Now())
}
