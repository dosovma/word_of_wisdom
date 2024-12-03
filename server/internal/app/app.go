package app

import (
	"fmt"

	"server/internal/api/tcp"
	"server/internal/config"
	"server/internal/service"
	"server/internal/storage"
)

func Run() error {
	cfg, err := config.NewServer()
	if err != nil {
		return err
	}

	tokenStorage := storage.NewTokenStorage()
	quoteStorage := storage.NewQuoteStorage()

	s := service.New(quoteStorage, tokenStorage)

	handler := tcp.NewHandler(s, tokenStorage)

	tcpServer := tcp.NewServer(cfg.Host, cfg.Port, handler)
	fmt.Println("tcp init")

	return tcpServer.Serve()
}
