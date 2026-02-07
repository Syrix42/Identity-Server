package application 

import (
	"context"
	"time"

	"github.com/alireza/identity/internal/domain"
	"github.com/alireza/identity/internal/infra/tokens"
	"github.com/google/uuid"
)

type LoginSerivce struct{
	Repo domain.UserRepo
	Comparer Comparer
}

func NewLoginService(repo domain.UserRepo , comparer Comparer ) *LoginSerivce{
	return &LoginSerivce{
		Repo: repo,
		Comparer: comparer,
	}
}


func (l *LoginSerivce) Login(ctx  context.Context ,UserName , Password string ) (string , string , error){
	existing , _ := l.Repo.GetByName(ctx ,UserName)
	if existing == nil{
		return "" , "" , ErrUserNotFound
	}
	if err := existing.CanAuthenticate(); err!=nil{
		return "" , "" , ErrCanNotAuthenticate
	}

	err := l.Comparer.Compare(ctx , existing.Password , Password)
	if err!= nil{
		return  "" , "" , err
	}
	err = l.Repo.IncrementActiveSessions(existing.UserID)
	if err!= nil{
		return "" , "" , err
	}
	accessToken := domain.NewAcessToken(uuid.NewString(),existing.ID(), domain.Access ,time.Now().UTC(), existing.Role)
	recoveryToken := domain.NewRecoveryToken(uuid.NewString(),existing.ID(), domain.Recovery , time.Now().UTC())
	path := PathLoader()
	pk , err := LoadPrivateKey(path)
	if err!= nil{
		return  "" , "" , err
	}

	clientAccessToken , err := tokens.IssueAccessToken(ctx , *accessToken , pk)
	if err!= nil{
		return "" , "" , err
	}

	clientRecoveryToken , err := tokens.IssueRecoveryToken(ctx , *recoveryToken , pk)
	if err!= nil{
		return "" , "" , err
	}



	return clientAccessToken , clientRecoveryToken , nil
	} 