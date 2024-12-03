package storage

import (
	"github.com/google/uuid"
	"server/internal/service/entity"
)

type TokenStorage map[uuid.UUID]entity.Token

func NewTokenStorage() TokenStorage {
	s := make(map[uuid.UUID]entity.Token)

	return s
}

func (s TokenStorage) Token(tokenID uuid.UUID) entity.Token {
	return entity.Token{}
}

func (s TokenStorage) Store(entity.Token) {}
