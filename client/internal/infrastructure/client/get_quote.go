package client

import (
	"client/pkg/tcp"
)

func (c Client) GetQuote(token string) (string, error) {
	date := []string{TOKEN + token}
	if err := c.m.Write(CMD_QUOTE, date); err != nil {
		return "", err
	}

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}

	quote, err := tcp.GetDataByHeader(QUOTE, messages)
	if err != nil {
		return "", err
	}

	return quote, nil
}
