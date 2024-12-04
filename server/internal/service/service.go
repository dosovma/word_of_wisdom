package service

import (
	"server/pkg/logger"
)

type Service struct {
	quoteStorage QuoteStorage
	tokenStorage TokenStorage
	log          logger.Logger
}

func New(quoteStorage QuoteStorage, tokenStorage TokenStorage, log logger.Logger) *Service {
	return &Service{
		quoteStorage: quoteStorage,
		tokenStorage: tokenStorage,
		log:          log,
	}
}
