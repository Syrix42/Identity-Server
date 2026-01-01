package domain

import "context"

type UserRepo interface {
	Save(ctx context.Context , u User) error
	GetByName(ctx context.Context , name string) (*User, error)
	GetById(ctx context.Context , id string)(*User , error)
	
}