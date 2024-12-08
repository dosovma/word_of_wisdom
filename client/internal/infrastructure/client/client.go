package client

import (
	"net"

	"client/pkg/logger"
)

//go:generate mockgen -destination=./mock/messenger.go -package=mock . ConnectionReaderWriter
type ConnectionReaderWriter interface {
	Write(conn net.Conn, messages []string) error
	Read(conn net.Conn) ([]string, error)
}

type Client struct {
	connection net.Conn
	connRW     ConnectionReaderWriter
	logger     logger.Logger
}

func NewTCPClient(conn net.Conn, connRW ConnectionReaderWriter, log logger.Logger) *Client {
	return &Client{
		connection: conn,
		connRW:     connRW,
		logger:     log,
	}
}
