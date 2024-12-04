package tcp

import (
	"net"

	"server/pkg/logger"
)

const (
	SERVER_PORT = ":9000"
)

type Server struct {
	host    string
	port    string
	handler *Handler
	log     logger.Logger
}

func NewServer(host string, port string, handler *Handler, logger logger.Logger) *Server {
	return &Server{
		host:    host,
		port:    port,
		handler: handler,
		log:     logger,
	}
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		s.log.Println(err)
		return err
	}
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			s.log.Printf("failed to close listener: %s\n", err)
		}
	}(listener)

	s.log.Printf("listener init, address: %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.log.Printf("failed to accept connection: %s", err)

			return err
		}
		s.log.Println("connection accepted")

		go s.handler.Handle(conn)
	}
}
