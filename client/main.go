package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
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

	requestTime := time.Now().Unix()
	data := []byte(fmt.Sprintf("%s::%d::%d\n", "/get_task", 1236, requestTime))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println("task requested")

	reader := bufio.NewReader(conn)

	task := ""
	for {
		task, err = reader.ReadString('\n')
		if err != nil {
			// log error
			return
		}
		fmt.Println(task)

		break
	}
	fmt.Println("task got")

	solution := ""
	if strings.Trim(task, " ") != "" {
		solution, err = solve(task)
		if err != nil {
			fmt.Println(err) // TODO запросить еще одну таску

			return
		}
	}

	data = []byte(fmt.Sprintf("%s::%s\n", "/solution", solution))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("solution sent")

	response := ""
	for {
		response, err = reader.ReadString('\n')
		if err != nil {
			// log error
			return
		}
		fmt.Println(response)

		break
	}

	<-time.After(10 * time.Second)
}
