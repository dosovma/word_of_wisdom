package app

import (
	"fmt"
	"log"
	"net"

	"client/internal/infrastructure/client"
	"client/internal/service"
)

const (
	SERVER_PORT = ":9000"
)

func Run() error {
	conn, err := net.Dial("tcp", SERVER_PORT)
	if err != nil {
		fmt.Println("dial error:", err)
		return err
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close connection: %s\n", err)
		}
	}(conn)

	c := client.NewTCPClient(conn)

	s := service.NewService(c)

	quote, err := s.Quote()
	if err != nil {
		return err
	}

	fmt.Println(quote)

	return nil
}
