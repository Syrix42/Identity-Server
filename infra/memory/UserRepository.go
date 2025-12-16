package memory

import (
	"errors"
	"fmt"

	"github.com/alireza/identity/domain"
)

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

func (r *InMemoryRepository) GetByName(user domain.User) (*domain.User, error) {
	for _, u := range r.database {
		if u.Name == user.Name {
			return &domain.User{
				Name:     u.Name,
				Password: u.Password,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}