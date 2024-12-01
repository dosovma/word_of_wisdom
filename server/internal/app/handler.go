package app

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
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
				return
			}

			time.Sleep(time.Second)
		}
		fmt.Printf("received: %s", bytes)

		request := strings.Split(string(bytes), "::")

		switch {
		case strings.HasPrefix(request[0], "/get_task"):
			fmt.Println("task request got")
			reqID, err := strconv.ParseInt(request[1], 10, 64)
			if err != nil {
				return // invalid response format
			}

			reqTime, err := strconv.ParseInt(request[2], 10, 64)
			if err != nil {
				return // invalid response format
			}

			data := []byte(Task(reqID, reqTime))
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println("task sent")
		case strings.HasPrefix(request[0], "/solution"):
			if isGranted := Validate(request[1]); isGranted {
			}
			_, err = conn.Write([]byte("access granted\n"))
			if err != nil {
				//log error
				return
			}
		}
	}
}
