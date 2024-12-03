package entity

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID         uuid.UUID
	ExpiryDate time.Time
}

type Request struct {
	ID        int64
	CreatedAt int64
}

type Token uuid.UUID

func GenerateToken() uuid.UUID {
	return uuid.New()
}
