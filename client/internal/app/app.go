package app

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"client/internal/infrastructure/client"
	"client/internal/service"
	"client/internal/service/solver"
	"client/pkg/tcp"
)

const (
	ServerPort = ":9000"
)

func Run() error {
	logger := log.New(os.Stdout, "client:", log.LstdFlags)

	conn, err := net.Dial("tcp", ServerPort)
	if err != nil {
		logger.Printf("failed to dial server: %s", err)

		return err
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close connection: %s", err)
		}
		log.Println("connection closed")
	}(conn)

	m := tcp.NewMessenger(logger, client.MessageStart, client.MessageEnd, client.MessageSizeLimit)

	c := client.NewTCPClient(conn, m, logger)
	slvr := solver.New(logger)

	return testRequest(service.NewService(c, slvr))
}

func testRequest(s service.IService) error {
	id := rand.Int63n(100000)
	t := time.Now().Unix()

	quote, err := s.Quote(id, t)
	if err != nil {
		return err
	}

	fmt.Println(quote)

	return nil
}
