package storage

import (
	"errors"
	"server/internal/service/entity"
	"time"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("token not found")

type TokenStorage map[uuid.UUID]int64

func NewTokenStorage() TokenStorage {
	return make(map[uuid.UUID]int64)
}

func (s TokenStorage) Token(tokenID uuid.UUID) (*entity.Token, error) {
	expTime, ok := s[tokenID]
	if !ok {
		return nil, ErrNotFound
	}

	return &entity.Token{
		ID:         tokenID,
		ExpiryDate: time.Unix(expTime, 0),
	}, nil
}

func (s TokenStorage) Store(token entity.Token) {
	s[token.ID] = token.ExpiryDate.Unix()
}
