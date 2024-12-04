package client

import (
	"client/pkg/tcp"
)

func (c Client) GetChallenge(requestID, requestTime string) (string, error) {
	date := []string{REQUEST_ID + requestID, REQUEST_TIME + requestTime}
	if err := c.m.Write(CMD_TOKEN, date); err != nil {
		return "", err
	}

	messages, err := c.m.Read()
	if err != nil {
		return "", err
	}

	challenge, err := tcp.GetDataByHeader(CHALLENGE, messages)
	if err != nil {
		return "", err
	}

	return challenge, nil
}
