package client

import (
	"client/pkg/tcp"
)

func (c *Client) GetChallenge(requestID, requestTime string) (string, error) {
	date := []string{COMMAND + CMD_TOKEN, REQUEST_ID + requestID, REQUEST_TIME + requestTime}
	if err := c.m.Write(date); err != nil {
		return "", err
	}
	c.l.Println("token requested")

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}
	c.l.Println("challenge message received")

	challenge, err := tcp.GetDataByHeader(CHALLENGE, messages)
	if err != nil {
		return "", err
	}
	c.l.Println("challenge found")

	return challenge, nil
}
