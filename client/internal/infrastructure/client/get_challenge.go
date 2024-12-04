package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetChallenge(requestID, requestTime string) (string, error) {
	data := []string{COMMAND + CMD_TOKEN, REQUEST_ID + requestID, REQUEST_TIME + requestTime}
	if err := c.messenger.Write(c.connection, data); err != nil {
		return "", err
	}
	c.logger.Println("token requested")

	messages, err := c.messenger.Read(c.connection)
	if err != nil {
		return "", err
	}
	c.logger.Println("challenge message received")

	challenge, err := tcp.GetDataByHeader(CHALLENGE, messages)
	if err != nil {
		return "", err
	}
	c.logger.Println("challenge found")

	return challenge, nil
}
