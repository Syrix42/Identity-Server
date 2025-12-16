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

	existing , _ := s.repo.GetByName(name)
	if existing != nil{
		return  ErrUserAlreadyExists
	}
	user := domain.NewUser(name , password)
	user.Password , _ = s.hasher.Hash(user.Password)
	err :=  s.repo.Save(user)
	if err != nil{
		return  err
	}
	return  nil
}

