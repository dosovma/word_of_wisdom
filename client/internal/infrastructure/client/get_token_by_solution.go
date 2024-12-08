package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetTokenBySolution(solution string) (string, error) {
	data := []string{Command + CmdSolution, Solution + solution}
	if err := c.connRW.Write(c.connection, data); err != nil {
		return "", err
	}

	c.logger.Println("solution sent")

	messages, err := c.connRW.Read(c.connection)
	if err != nil {
		return "", err
	}

	c.logger.Println("token message received")

	token, err := tcp.GetDataByHeader(Token, messages)
	if err != nil {
		return "", err
	}

	c.logger.Println("token found")

	return token, nil
}
