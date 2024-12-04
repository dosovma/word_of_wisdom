package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetQuote(token string) (string, error) {
	date := []string{COMMAND + CMD_QUOTE, TOKEN + token}
	if err := c.m.Write(date); err != nil {
		return "", err
	}
	c.l.Println("quote requested")

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}
	c.l.Println("quote message received")

	quote, err := tcp.GetDataByHeader(QUOTE, messages)
	if err != nil {
		return "", err
	}
	c.l.Println("quote found")

	return quote, nil
}
