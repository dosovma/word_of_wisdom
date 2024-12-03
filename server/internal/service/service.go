package service

import (
	"github.com/google/uuid"
	"server/internal/service/entity"
)

type QuoteStorage interface {
	Quote() (string, error)
}

type TokenStorage interface {
	Token(tokenID uuid.UUID) entity.Token
	Store(entity.Token)
}

type IService interface {
	Quote() (string, error)
	Challenge(request entity.Request) string
	Validate(solution string) bool
	Token() uuid.UUID
}

type Service struct {
	quoteStorage QuoteStorage
	tokenStorage TokenStorage
}

func New(quoteStorage QuoteStorage, tokenStorage TokenStorage) *Service {
	return &Service{
		quoteStorage: quoteStorage,
		tokenStorage: tokenStorage,
	}
}
