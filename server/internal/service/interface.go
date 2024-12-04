package service

import (
	"github.com/google/uuid"

	"server/internal/service/entity"
)

// input interface

type IService interface {
	Quote() (string, error)
	Challenge(request entity.Request) string
	Validate(solution string) bool
	Token() uuid.UUID
}

// output interfaces

//go:generate mockgen -destination=./mock/quote_storage.go -package=mock . QuoteStorage
type QuoteStorage interface {
	Quote() (string, error)
}

//go:generate mockgen -destination=./mock/token_storage.go -package=mock . TokenStorage
type TokenStorage interface {
	Token(tokenID uuid.UUID) (*entity.Token, error)
	Store(entity.Token)
}
