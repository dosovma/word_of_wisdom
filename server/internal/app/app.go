package app

import (
	"context"
	"log"
	"os"

	"server/internal/api/tcp"
	"server/internal/config"
	"server/internal/service"
	"server/internal/storage"
	messenger "server/pkg/tcp"
)

func Run() error {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Println(context.Background(), panicErr)
		}
	}()

	logger := log.New(os.Stdout, "server:", log.LstdFlags)

	cfg, err := config.NewServer()
	if err != nil {
		return err
	}

	tokenStorage := storage.NewTokenStorage()

	quoteStorage, err := storage.NewQuoteStorage()
	if err != nil {
		logger.Printf("failed to init quote storage: %s", err)
		return err
	}

	s := service.New(quoteStorage, tokenStorage, logger)

	connRW := messenger.NewConnectionRW(logger, tcp.MessageStart, tcp.MessageEnd, tcp.MessageSizeLimit)

	handler := tcp.NewHandler(s, tokenStorage, connRW, logger)

	tcpServer := tcp.NewServer(cfg.Host, cfg.Port, cfg.Timeout, handler, logger)
	logger.Println("server init")

	return tcpServer.Serve()
}
