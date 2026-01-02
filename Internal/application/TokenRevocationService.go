package application

import (
	"github.com/alireza/identity/internal/domain"
	"errors"
)


type TokenRevocationService struct{
	repo domain.TokenRepository
}

func NewTokenRevocationService(repo domain.TokenRepository) TokenRevocationService {
	return TokenRevocationService{
		repo: repo,
	}
}

func (t *TokenRevocationService) RevokeToken(refreshToken string) error {
	
}