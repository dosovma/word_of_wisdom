package app

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Handler struct{}

func (h *Handler) Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			// log
		}
	}(conn)

	reader := bufio.NewReader(conn)
	fmt.Println("got connection")

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("received: %s", bytes)

		_, err = conn.Write([]byte("ack\n"))
		if err != nil {
			//log error
			return
		}

		err = conn.Close()
		if err != nil {
			fmt.Printf("failed to close connection %s\n", err)
			return
		}
	}
}
