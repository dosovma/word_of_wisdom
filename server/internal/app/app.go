package app

import (
	"log"
	"os"

	"server/internal/api/tcp"
	"server/internal/config"
	"server/internal/service"
	"server/internal/storage"
	messenger "server/pkg/tcp"
)

func Run() error {
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

	m := messenger.NewMessenger(logger, tcp.MessageStart, tcp.MessageEnd, tcp.MessageSizeLimit)

	handler := tcp.NewHandler(s, tokenStorage, m, logger)

	tcpServer := tcp.NewServer(cfg.Host, cfg.Port, handler, logger)
	logger.Println("server init")

	return tcpServer.Serve()
}
