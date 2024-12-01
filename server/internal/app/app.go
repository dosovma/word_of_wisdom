package app

import (
	"fmt"

	"server/internal/config"
)

func Run() error {
	fmt.Println("Server starts")
	cfg, err := config.NewServer()
	if err != nil {
		return err
	}

	handler := Handler{}

	tcpServer := New(cfg.Host, cfg.Port, handler)
	fmt.Println("tcp init")

	return tcpServer.Serve()
}
