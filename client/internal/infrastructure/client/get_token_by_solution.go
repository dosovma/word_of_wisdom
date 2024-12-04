package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetTokenBySolution(solution string) (string, error) {
	date := []string{COMMAND + CMD_SOLUTION, SOLUTION + solution}
	if err := c.m.Write(date); err != nil {
		return "", err
	}
	c.l.Println("solution sent")

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}
	c.l.Println("token message received")

	token, err := tcp.GetDataByHeader(TOKEN, messages)
	if err != nil {
		return "", err
	}
	c.l.Println("token found")

	return token, nil
}
