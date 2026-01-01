package application 

import (
	"context"

	"github.com/alireza/identity/internal/domain"
	"github.com/google/uuid"
)


type UserService struct{
	repo domain.UserRepo
	hasher PasswordHasher
}


func NewUserService(repo domain.UserRepo , hasher PasswordHasher) *UserService{
	return  &UserService{repo: repo, hasher: hasher}
}

func (s *UserService) Register(ctx context.Context, name , password string , role string) error {
	existing , _ :=  s.repo.GetByName(ctx ,name)
	if existing != nil{
		return  ErrUserAlreadyExists
	}
	user := domain.NewUser(uuid.NewString(), name , password , role)
	user.Password , _ = s.hasher.Hash(ctx ,user.Password)
	err :=  s.repo.Save(ctx , user)
	if err != nil{
		return  err
	}
	return  nil
}

