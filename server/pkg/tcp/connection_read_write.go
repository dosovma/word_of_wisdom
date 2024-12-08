package tcp

import (
	"bufio"
	"errors"
	"io"
	"net"
	"strings"

	"server/pkg/logger"
)

var (
	ErrLimitExceeded = errors.New("message limit exceeded")
	ErrNotFound      = errors.New("header not found")
)

type ConnRW struct {
	logger       logger.Logger
	messageStart string
	messageEnd   string
	messageLimit int
}

func NewConnectionRW(logger logger.Logger, messageStart string, messageEnd string, messageLimit int) *ConnRW {
	return &ConnRW{
		logger:       logger,
		messageStart: messageStart,
		messageEnd:   messageEnd,
		messageLimit: messageLimit,
	}
}

func (crw *ConnRW) Write(conn net.Conn, messages []string) error {
	message := make([]string, 0, len(messages)+2) //nolint:mnd
	message = append(message, crw.messageStart)
	message = append(message, messages...)
	message = append(message, crw.messageEnd)

	for _, msg := range message {
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			crw.logger.Printf("failed to write message: %s: err: %s", msg, err)

			return err
		}
	}

	return nil
}

func (crw *ConnRW) Read(conn net.Conn) ([]string, error) {
	messageSize := 0
	connReader := bufio.NewReader(conn)

	message := make([]string, 0, 1)

	for {
		str, err := connReader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				crw.logger.Printf("invalid message format: not found END message: %s", err)

				return message, nil
			}

			crw.logger.Printf("failed to read message: %s", err)

			return nil, err
		}

		msg, _ := strings.CutSuffix(str, "\n")

		messageSize += len(msg)
		if messageSize >= crw.messageLimit {
			crw.logger.Printf("message limit exceeded: %d", messageSize)

			return nil, ErrLimitExceeded
		}

		if crw.isPayload(msg) {
			message = append(message, msg)
		}

		if msg == crw.messageEnd {
			return message, nil
		}
	}
}

func (crw *ConnRW) isPayload(msg string) bool {
	return msg != crw.messageStart && msg != crw.messageEnd
}

func GetDataByHeader(header string, messages []string) (string, error) {
	for _, str := range messages {
		if strings.HasPrefix(str, header) {
			data, _ := strings.CutPrefix(str, header)

			return data, nil
		}
	}

	return "", ErrNotFound
}
