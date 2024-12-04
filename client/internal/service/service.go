package service

import (
	"client/internal/infrastructure/client"
)

type IService interface {
	Quote() (string, error)
}

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) *Service {
	return &Service{client: client}
}
