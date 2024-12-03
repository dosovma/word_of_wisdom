package tcp

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNotFound = errors.New("header not found")

func GetDataByHeader(header string, messages []string) (string, error) {
	for _, str := range messages {
		if strings.HasPrefix(str, header) {
			data, _ := strings.CutPrefix(str, header)

			return data, nil
		}
	}

	return "", ErrNotFound
}

func AddDataToHeader(header string, messages []string) []string {
	for _, msg := range messages {
		messages = append(messages, fmt.Sprintf("%s%s\n", header, msg))
	}

	return messages
}
