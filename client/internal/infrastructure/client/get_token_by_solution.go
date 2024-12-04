package client

import (
	"client/pkg/tcp"
)

func (c Client) GetTokenBySolution(solution string) (string, error) {
	date := []string{SOLUTION + solution}
	if err := c.m.Write(CMD_SOLUTION, date); err != nil {
		return "", err
	}

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}

	token, err := tcp.GetDataByHeader(TOKEN, messages)
	if err != nil {
		return "", err
	}

	return token, nil
}
