package memory

import (
	"context"
	"errors"
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



 func(r *InMemoryRepository) Save(ctx context.Context ,user domain.User) error {
	
	r.mutex.Lock()
	r.database = append(r.database, user)
	defer r.mutex.Unlock()
	return  nil	
}

func (r *InMemoryRepository) GetByName(ctx context.Context , name string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, u := range r.database {
		if u.UserName == name {
			user:= domain.NewUser(u.ID(),u.UserName , u.Password ,u.Role)
			return &user ,nil
		}
	}

	return nil, ErrUserNotFound
}

func (r *InMemoryRepository)GetById(ctx context.Context , id  domain.UserID ) (*domain.User , error){
	return nil , nil
}