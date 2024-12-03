package service

import (
	"time"

	"github.com/google/uuid"
	"server/internal/service/entity"
)

const EXPIRY_DURATION = 1 * 60 * 60 // 1 час

func (s *Service) Token() uuid.UUID {
	t := entity.Token{
		ID:         uuid.New(),
		ExpiryDate: time.Now().Add(time.Second * EXPIRY_DURATION),
	}

	s.tokenStorage.Store(t)

	return t.ID
}
