package client

import (
	"net"

	"client/pkg/logger"
	"client/pkg/tcp"
)

type messenger interface {
	Write(messages []string) error
	Read() ([]string, error)
}

type Client struct {
	conn net.Conn
	m    messenger
	l    logger.Logger
}

func NewTCPClient(conn net.Conn, l logger.Logger) *Client {
	return &Client{
		conn: conn,
		m:    tcp.NewMessenger(conn, MESSAGE_START, MESSAGE_END, MESSAGE_SIZE_LIMIT),
		l:    l,
	}
}
