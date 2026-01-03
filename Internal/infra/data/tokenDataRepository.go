package data



import (
	"context"
	database "github.com/alireza/identity/Database"
	"time"
)




type TokenDb struct {
	Id   string  `db:"jit"`
	ExpiresAt time.Time `db:"expires_at"`
	AddedAt time.Time `db:"added_time"`
}
type TokenDataRepository struct {}

func NewTokenDataRepository() *TokenDataRepository {
	return &TokenDataRepository{
	}
}
func (t *TokenDataRepository) RevokeToken(ctx context.Context, jit string, expiresAt time.Time) error {
	db := database.GetDB()
	addedAt := time.Now().UTC()
	query := "INSERT INTO revoked_tokens (jit, expires_at, added_time) VALUES (:jit, :expires_at, :added_time)"
	_, err := db.NamedExecContext(ctx, query, TokenDb{
		Id:        jit,
		ExpiresAt: expiresAt,	
		AddedAt:   addedAt,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenDataRepository) IsTokenRevoked(ctx context.Context, jit string) (bool , error)  {
	db := database.GetDB()
	var count int
	query := "SELECT COUNT(*) FROM revoked_tokens WHERE jit = ?" 
	err := db.GetContext(ctx, &count, query, jit)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true , nil
	}
	return false ,  nil
}
