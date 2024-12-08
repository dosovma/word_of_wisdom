package tcp

import (
	"context"
	"net"
	"time"

	"server/pkg/logger"
)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
}

type Server struct {
	host    string
	port    string
	timeout int
	handler Handler
	log     logger.Logger
}

func NewServer(host string, port string, timeout int, handler Handler, logger logger.Logger) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: timeout,
		handler: handler,
		log:     logger,
	}
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp4", s.host+s.port)
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
		}

		s.log.Println("connection accepted")

		go s.processConnection(conn)
	}
}

func (s *Server) processConnection(conn net.Conn) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*time.Duration(s.timeout))
	defer cancel()

	s.handler.Handle(ctx, conn)
}
