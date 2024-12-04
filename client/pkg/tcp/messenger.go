package tcp

import (
	"bufio"
	"errors"
	"net"
)

type Messenger struct {
	conn         net.Conn
	messageStart string
	messageEnd   string
	messageLimit int
}

func NewMessenger(conn net.Conn, messageStart string, messageEnd string, messageLimit int) *Messenger {
	return &Messenger{conn: conn, messageStart: messageStart, messageEnd: messageEnd, messageLimit: messageLimit}
}

func (m *Messenger) Write(header string, messages []string) error {
	message := make([]string, 0, len(messages)+2)
	message = append(message, m.messageStart)
	message = append(message, AddDataToHeader(header, messages)...)
	message = append(message, m.messageEnd)

	for _, msg := range message {
		_, err := m.conn.Write([]byte(msg))
		if err != nil {
			return err // TODO
		}
	}

	return nil
}

func (m *Messenger) Read() ([]string, error) {
	messageSize := 0
	connReader := bufio.NewReader(m.conn)

	message := make([]string, 0, 1)

	for {
		str, err := connReader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		messageSize += len(str)
		if messageSize >= m.messageLimit {
			return nil, errors.New("message limit exceeded")
		}

		if m.isPayload(str) {
			message = append(message, str)
		}

		if str == m.messageEnd {
			return message, nil
		}
	}
}

func (m *Messenger) isPayload(msg string) bool {
	return msg != m.messageStart && msg != m.messageEnd
}
