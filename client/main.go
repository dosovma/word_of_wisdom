package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	SERVER_PORT = ":9000"
)

func main() {
	conn, err := net.Dial("tcp", SERVER_PORT)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	data := []byte("Hello, Server!\n")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("message sent")

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// log error
			return
		}
		fmt.Println(message)
		break
	}

	<-time.After(10 * time.Second)
}
