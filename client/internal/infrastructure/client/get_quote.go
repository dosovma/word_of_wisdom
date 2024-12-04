package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetQuote(token string) (string, error) {
	date := []string{Command + CmdQuote, Token + token}
	if err := c.messenger.Write(c.connection, date); err != nil {
		return "", err
	}
	c.logger.Println("quote requested")

	messages, err := c.messenger.Read(c.connection)
	if err != nil {
		return "", err
	}
	c.logger.Println("quote message received")

	quote, err := tcp.GetDataByHeader(Quote, messages)
	if err != nil {
		return "", err
	}
	c.logger.Println("quote found")

	return quote, nil
}
