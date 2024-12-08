package tcp

import (
	"context"
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

type connectionReaderWriter interface {
	Write(conn net.Conn, messages []string) error
	Read(conn net.Conn) ([]string, error)
}

type HandlerTCP struct {
	service      service.IService
	connRW       connectionReaderWriter
	tokenStorage tokenStorage
	log          logger.Logger
}

func NewHandler(service service.IService, storage tokenStorage, connRW connectionReaderWriter, log logger.Logger) *HandlerTCP {
	return &HandlerTCP{
		service:      service,
		tokenStorage: storage,
		log:          log,
		connRW:       connRW,
	}
}

func (h *HandlerTCP) Handle(ctx context.Context, conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			h.log.Printf("failed to close connection: %s\n", err)
		}

		h.log.Println("connection closed")
	}(conn)

	done := make(chan bool, 1)
	defer close(done)

	for {
		select {
		case <-ctx.Done():
			h.log.Printf("timeout connection exceeded")

			return
		case <-done:
			return
		default:
			h.processMessage(done, conn)
		}
	}
}

func (h *HandlerTCP) processMessage(done chan bool, conn net.Conn) {
	msg, err := h.connRW.Read(conn)
	if err != nil {
		h.log.Println("failed to read message")

		return // TODO add error handling
	}

	cmd, err := tcp.GetDataByHeader(Command, msg)
	if err != nil {
		h.log.Println("failed to read command")

		return
	}

	switch cmd {
	case CmdToken:
		h.processToken(conn, msg)
	case CmdSolution:
		h.processSolution(conn, msg)
	case CmdQuote:
		h.processQuote(conn, msg)

		h.log.Println("quote sent")

		done <- true

		return
	default:
		h.log.Println("unknown command header")
	}
}

func (h *HandlerTCP) processToken(conn net.Conn, msg []string) {
	request, err := extractRequest(msg)
	if err != nil {
		h.log.Println("failed to get request header")

		return
	}

	challenge := h.service.Challenge(*request)

	if err = h.connRW.Write(conn, []string{Challenge + challenge}); err != nil {
		h.log.Println("failed to send challenge")

		return
	}
}

func (h *HandlerTCP) processSolution(conn net.Conn, msg []string) {
	solution, err := tcp.GetDataByHeader(Solution, msg)
	if err != nil {
		h.log.Println("failed to get solution header")

		return
	}

	if isGranted := h.service.Validate(solution); !isGranted {
		if err = h.connRW.Write(conn, []string{Access + "Reject"}); err != nil {
			return
		}

		return
	}

	token := h.service.Token()

	if err = h.connRW.Write(conn, []string{Token + token.String()}); err != nil {
		h.log.Println("failed to send token")

		return
	}
}

func (h *HandlerTCP) processQuote(conn net.Conn, msg []string) {
	if h.isAuth(msg) {
		q, err := h.service.Quote()
		if err != nil {
			h.log.Println("failed to get quote from service")

			return
		}

		if err = h.connRW.Write(conn, []string{Quote + q}); err != nil {
			h.log.Println("failed to send quote")

			return
		}
	}

	if err := h.connRW.Write(conn, []string{Access + "Reject"}); err != nil {
		return
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

func (h *HandlerTCP) isAuth(messages []string) bool {
	tokenStr, err := tcp.GetDataByHeader(Token, messages)
	if err != nil {
		return false
	}

	tokenID, err := uuid.Parse(tokenStr)
	if err != nil {
		return false
	}

	token, err := h.tokenStorage.Token(tokenID)
	if err != nil {
		h.log.Printf("failed to get token: %s", err)

		return false
	}

	return token.ExpiryDate.After(time.Now())
}
