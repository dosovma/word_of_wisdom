package service

import "errors"

var ErrQuoteNotFound = errors.New("quote not found")

func (s *Service) Quote() (string, error) {
	q, err := s.storage.Quote()
	if err != nil {
		return "", ErrQuoteNotFound
	}

	return q, nil
}
