package client

import (
	"client/pkg/logger"
	"net"
)

//go:generate mockgen -destination=./mock/messenger.go -package=mock . Messenger
type Messenger interface {
	Write(conn net.Conn, messages []string) error
	Read(conn net.Conn) ([]string, error)
}

type Client struct {
	connection net.Conn
	messenger  Messenger
	logger     logger.Logger
}

func NewTCPClient(conn net.Conn, m Messenger, log logger.Logger) *Client {
	return &Client{
		connection: conn,
		messenger:  m,
		logger:     log,
	}
}
