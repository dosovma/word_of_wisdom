package storage

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	growOutPath = "internal/storage/config/fetcher/files/listing_grow_out.yaml"
)

var ErrEmptyStorage = errors.New("storage is empty")

type QuoteStorage map[string]struct{}

func NewQuoteStorage() (QuoteStorage, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get root path, details %w", err)
	}

	path := fmt.Sprintf("%s/%s", root, growOutPath)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file with path %s, details %w", path, err)
	}

	reader := bufio.NewReader(file)

	storage := make(map[string]struct{})

	for {
		quote, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		storage[quote] = struct{}{}
	}

	return storage, nil
}

func (s QuoteStorage) Quote() (string, error) {
	for q := range s {
		return q, nil
	}

	return "", ErrEmptyStorage
}
