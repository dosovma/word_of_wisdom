package app

import (
	"fmt"

	"server/internal/api/tcp"
	"server/internal/config"
)

func Run() error {
	cfg, err := config.NewServer()
	if err != nil {
		return err
	}

	handler := tcp.Handler{}

	tcpServer := tcp.NewServer(cfg.Host, cfg.Port, handler)
	fmt.Println("tcp init")

	return tcpServer.Serve()
}
