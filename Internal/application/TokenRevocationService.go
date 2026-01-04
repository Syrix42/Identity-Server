package application

import (
	"context"
	"time"

	"log"

	"github.com/alireza/identity/internal/domain"
	"github.com/alireza/identity/internal/infra/tokens"
	"github.com/google/uuid"
	"github.com/alireza/identity/internal/infra/revokedtoken"
)


type TokenRevocationService struct{
	repo domain.TokenRepository
	Userrepo domain.UserRepo
}

func NewTokenRevocationService(repo domain.TokenRepository, userRepo domain.UserRepo) TokenRevocationService {
	return TokenRevocationService{
		repo: repo,
		Userrepo: userRepo,
	}
}

func (t *TokenRevocationService) RevokeToken(ctx context.Context, refreshToken string) (string ,string ,error) {
    publickey, err := LoadPublicKey()
	if err != nil || publickey == nil {
    log.Fatal("failed to load public key:", err)
}
	claims , err := tokens.ValidateRecoveryToken(ctx , refreshToken, publickey)
	if err != nil {
		return "", "", err
	}
	TableDate , err := revokedtoken.EnsureRevokedTokenTable(ctx , claims.IssuedAt.Time)
	if err!=nil{
		log.Fatal("Could not form the table")
	}
	if err!=nil{
		return "" , "" , err
	}
	isRevoked, err := t.repo.IsTokenRevoked(ctx ,claims.ID, TableDate)
	if err != nil{
		return "" , "" , err	
	}
	if isRevoked {
		return "" , "" , ErrTokenAlreadyRevoked
	}

	t.repo.RevokeToken(ctx, claims.ID , TableDate ,claims.ExpiresAt.Time.UTC())
	path := PathLoader()
	privateKey , err := LoadPrivateKey(path)
	if err!= nil{
		return "","",err
	}
	User , err:= t.Userrepo.GetById(ctx , claims.Subject)
	if err!= nil{
		return "" , "" ,err
	}
	Accesstoken := domain.NewAcessToken(uuid.NewString(),claims.Subject,domain.Access,time.Now().UTC(), User.Role)
	RecoveryToken := domain.NewRecoveryToken(uuid.NewString(),claims.Subject,domain.Recovery,time.Now().UTC())
	strRecoveryToken ,  err := tokens.IssueRecoveryToken(ctx , *RecoveryToken , privateKey)
	strAccessToken , err := tokens.IssueAccessToken(ctx , *Accesstoken , privateKey)

	return strAccessToken , strRecoveryToken , nil
}