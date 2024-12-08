package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetChallenge(requestID, requestTime string) (string, error) {
	data := []string{Command + CmdToken, RequestID + requestID, RequestTime + requestTime}
	if err := c.connRW.Write(c.connection, data); err != nil {
		return "", err
	}

	c.logger.Println("token requested")

	messages, err := c.connRW.Read(c.connection)
	if err != nil {
		return "", err
	}

	c.logger.Println("challenge message received")

	challenge, err := tcp.GetDataByHeader(Challenge, messages)
	if err != nil {
		return "", err
	}

	c.logger.Println("challenge found")

	return challenge, nil
}
