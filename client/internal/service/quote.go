package service

import (
	"strconv"
)

func (s *Service) Quote(requestID, requestTime int64) (string, error) {
	reqID := strconv.FormatInt(requestID, 10)
	reqTime := strconv.FormatInt(requestTime, 10)

	challenge, err := s.client.GetChallenge(reqID, reqTime)
	if err != nil {
		return "", err
	}

	solution, err := s.solver.Solve(challenge)
	if err != nil {
		return "", err // TODO repeat request
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
