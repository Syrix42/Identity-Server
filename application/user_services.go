package application


import (
	"github.com/alireza/identity/domain"
)

type UserService struct{
	repo domain.UserRepo
	hasher PasswordHasher
}


func NewUserService(repo domain.UserRepo , hasher PasswordHasher) *UserService{
	return  &UserService{repo: repo, hasher: hasher}
}

func (s *UserService) Register(name , password string) error {
	user := domain.NewUser(name , password)
	user.Password , _ = s.hasher.Hash(user.Password)
	return  s.repo.Save(user)
}

