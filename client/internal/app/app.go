package app

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"time"

	"client/internal/infrastructure/client"
	"client/internal/service"
	"client/internal/service/solver"
	"client/pkg/tcp"
)

const (
	ServerHost = "127.0.0.1"
	ServerPort = ":9000"
)

func Run() error {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Println(context.Background(), panicErr)
		}
	}()

	logger := log.New(os.Stdout, "client:", log.LstdFlags)

	conn, err := net.Dial("tcp4", ServerHost+ServerPort)
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
	quote, err := s.Quote(rand.Int64N(100500), time.Now().Unix()) //nolint:gosec,mnd
	if err != nil {
		return err
	}

	fmt.Println(quote)

	return nil
}
