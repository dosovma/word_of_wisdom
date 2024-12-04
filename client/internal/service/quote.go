package service

import (
	"math/rand"
	"strconv"
	"time"
)

func (s *Service) Quote() (string, error) {
	id := rand.Int63n(100000)
	t := time.Now().Unix()

	reqID := strconv.FormatInt(id, 10)
	reqTime := strconv.FormatInt(t, 10)

	challenge, err := s.client.GetChallenge(reqID, reqTime)
	if err != nil {
		return "", err
	}

	solution, err := solve(challenge)
	if err != nil {
		return "", err
	}

	token, err := s.client.GetTokenBySolution(solution)
	if err != nil {
		return "", err
	}

	quote, err := s.client.GetQuote(token)
	if err != nil {
		return "", err
	}

	return quote, nil
}
