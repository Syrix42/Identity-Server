package domain

import (
	"context"
	"time"
)
// I dident followed CQRS Princple here for simplicity because domain is not that rich

type TokenRepository interface {
	RevokeToken(ctx context.Context, ID string, Expiration time.Time) error
	IsTokenRevoked(ctx context.Context, ID string) (bool,error)
}