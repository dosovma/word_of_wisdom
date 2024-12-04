package service

import (
	"time"

	"github.com/google/uuid"

	"server/internal/service/entity"
)

const ExpiryTime = 1 * 60 * 60 // 1 час // TODO set by envs

func (s *Service) Token() uuid.UUID {
	t := entity.Token{
		ID:         uuid.New(),
		ExpiryDate: time.Now().Add(time.Second * ExpiryTime),
	}

	s.tokenStorage.Store(t)

	return t.ID
}
