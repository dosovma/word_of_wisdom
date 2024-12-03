package storage

import "errors"

func init() {
	// init storage
}

var ErrEmptyStorage = errors.New("storage is empty")

type QuoteStorage map[string]struct{}

func NewQuoteStorage() QuoteStorage {
	s := make(map[string]struct{})

	return s
}

func (s QuoteStorage) Quote() (string, error) {
	for q := range s {
		return q, nil
	}

	return "", ErrEmptyStorage
}
