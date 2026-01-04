package application

import (
	"context"
	"log"

	"github.com/alireza/identity/internal/domain"
	"github.com/alireza/identity/internal/infra/revokedtoken"
	"github.com/alireza/identity/internal/infra/tokens"
)

type LogOutService struct {
	userRepo domain.UserRepo
	tokenRepo domain.TokenRepository
}

func NewLogOutService(userRepo domain.UserRepo , tokenrepo domain.TokenRepository) *LogOutService {
	return &LogOutService{
		userRepo: userRepo,
		tokenRepo: tokenrepo,
	}
}


func (h *LogOutService) Logout(ctx context.Context , refreshToken string) error {
	publicKey ,  err := LoadPublicKey()

	if err!= nil{
		log.Fatal("Failed to Load the Public Key")
	}

	claims , err := tokens.ValidateRecoveryToken(ctx , refreshToken ,publicKey )
	if err!= nil {
		return  err
	}
	tableName, err := revokedtoken.EnsureRevokedTokenTable(ctx, claims.IssuedAt.Time)

	if err!= nil{
		return  err
	}
	isRevoked , err := h.tokenRepo.IsTokenRevoked(ctx , claims.ID ,tableName)
	if err != nil {
		return err
	}
	if isRevoked{
		return ErrInvalidToken
	}
	err = h.tokenRepo.RevokeToken(ctx , claims.ID ,tableName , claims.ExpiresAt.Time.UTC())
	if err!= nil{
		return err
	}

	err = h.userRepo.DecrementActiveSessions(claims.Subject)
	if err!= nil{
		return err
	}
	return nil
}


