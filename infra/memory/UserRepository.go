package memory

import (
	"errors"
	"fmt"

	"github.com/alireza/identity/domain"
)

var ErrUserNotFound = errors.New("user not found")


type InMemoryRepository struct{
	database []domain.User
}
func NewInMemoryRepository() *InMemoryRepository{
	return &InMemoryRepository{
		database : []domain.User{},
	}


}



 func(r *InMemoryRepository) Save(user domain.User) error {
	
	
	r.database = append(r.database, user)
	fmt.Println(user.Name , user.Password)
	return  nil
}

func (r *InMemoryRepository) GetByName(name string) (*domain.User, error) {
	for _, u := range r.database {
		if u.Name == name {
			return &domain.User{
				Name:     u.Name,
				Password: u.Password,
			}, nil
		}
	}
	return nil, ErrUserNotFound
}