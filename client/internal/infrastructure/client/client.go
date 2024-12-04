package client

import (
	"net"

	"client/pkg/tcp"
)

type messenger interface {
	Write(header string, messages []string) error
	Read() ([]string, error)
}

type Client struct {
	conn net.Conn
	m    messenger
}

func NewTCPClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
		m:    tcp.NewMessenger(conn, MESSAGE_START, MESSAGE_END, MESSAGE_SIZE_LIMIT),
	}
}
