package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
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

func (m *Messenger) Write(messages []string) error {
	message := make([]string, 0, len(messages)+2)
	message = append(message, m.messageStart)
	message = append(message, messages...)
	message = append(message, m.messageEnd)

	for _, msg := range message {
		_, err := m.conn.Write([]byte(msg + "\n"))
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
		fmt.Println("start reading")
		str, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read data")
			return nil, err
		}

		msg, _ := strings.CutSuffix(str, "\n")

		messageSize += len(msg)
		if messageSize >= m.messageLimit {
			fmt.Println("message limit exceeded")
			return nil, errors.New("message limit exceeded")
		}

		if m.isPayload(msg) {
			message = append(message, msg)
		}

		if msg == m.messageEnd {
			fmt.Println("finish reading")
			return message, nil
		}
	}
}

func (m *Messenger) isPayload(msg string) bool {
	return msg != m.messageStart && msg != m.messageEnd
}
