package application

import "context"

type PasswordHasher interface {
	Hash(ctx context.Context , plaintext string) (string, error)
}