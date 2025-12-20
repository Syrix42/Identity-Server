package application

import (

	"github.com/alireza/identity/domain"
)

type LoginSerivce struct{
	repo domain.UserRepo
	comparer Comparer
}

func NewLoginService(repo domain.UserRepo ,Comparer Comparer) *LoginSerivce{
	return  &LoginSerivce{
		repo:  repo,
		comparer: Comparer,
	}
}


func (l *LoginSerivce) Login(UserName , Password string ) error{
	existing , _ := l.repo.GetByName(UserName)
	if existing == nil{
		return  ErrUserNotFound
	}
	err:= l.comparer.Compare(existing.Password , Password)
	if err!= nil{
		return err
	}
	return nil
} 