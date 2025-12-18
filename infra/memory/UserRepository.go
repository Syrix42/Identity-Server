package memory

import (
	"errors"
	"fmt"
	"sync"
	"github.com/alireza/identity/domain"
)

var ErrUserNotFound = errors.New("user not found")


type InMemoryRepository struct{
	database []domain.User
	mutex sync.RWMutex
}
func NewInMemoryRepository() *InMemoryRepository{
	return &InMemoryRepository{
		database : []domain.User{},
	}


}



 func(r *InMemoryRepository) Save(user domain.User) error {
	
	r.mutex.Lock()
	r.database = append(r.database, user)
	fmt.Println(user.Name , user.Password)
	defer r.mutex.Unlock()
	return  nil
}

func (r *InMemoryRepository) GetByName(name string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
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