package app

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host    string
	port    string
	handler Handler
}

const (
	SERVER_PORT = ":9000"
)

func New(host string, port string, handler Handler) *Server {
	return &Server{
		host:    host,
		port:    port,
		handler: handler,
	}
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		return err
	}
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			log.Printf("failed to close listener: %s\n", err)
		}
	}(listener)

	fmt.Printf("listener init, address: %s\n", listener.Addr())

	for {
		fmt.Printf("run listening\n")
		conn, err := listener.Accept()
		fmt.Println("connection accepted")
		if err != nil {
			fmt.Println("failed to accept connection")
		}

		// ToDo context?
		go s.handler.Handle(conn)
	}
}
