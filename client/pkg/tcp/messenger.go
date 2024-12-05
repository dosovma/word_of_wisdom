package tcp

import (
	"bufio"
	"errors"
	"io"
	"net"
	"strings"

	"client/pkg/logger"
)

var (
	ErrLimitExceeded = errors.New("message limit exceeded")
	ErrNotFound      = errors.New("header not found")
)

type Messenger struct {
	logger       logger.Logger
	messageStart string
	messageEnd   string
	messageLimit int
}

func NewMessenger(logger logger.Logger, messageStart string, messageEnd string, messageLimit int) *Messenger {
	return &Messenger{
		logger:       logger,
		messageStart: messageStart,
		messageEnd:   messageEnd,
		messageLimit: messageLimit,
	}
}

func (m *Messenger) Write(conn net.Conn, messages []string) error {
	message := make([]string, 0, len(messages)+2) //nolint:mnd
	message = append(message, m.messageStart)
	message = append(message, messages...)
	message = append(message, m.messageEnd)

	for _, msg := range message {
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			m.logger.Printf("failed to write message: %s: err: %s", msg, err)

			return err
		}
	}

	return nil
}

func (m *Messenger) Read(conn net.Conn) ([]string, error) {
	messageSize := 0
	connReader := bufio.NewReader(conn)

	message := make([]string, 0, 1)

	for {
		str, err := connReader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				m.logger.Printf("invalid message format: not found END message: %s", err)

				return message, nil
			}

			m.logger.Printf("failed to read message: %s", err)

			return nil, err
		}

		msg, _ := strings.CutSuffix(str, "\n")

		messageSize += len(msg)
		if messageSize >= m.messageLimit {
			m.logger.Printf("message limit exceeded: %d", messageSize)

			return nil, ErrLimitExceeded
		}

		if m.isPayload(msg) {
			message = append(message, msg)
		}

		if msg == m.messageEnd {
			return message, nil
		}
	}
}

func (m *Messenger) isPayload(msg string) bool {
	return msg != m.messageStart && msg != m.messageEnd
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
