package data



import (
	"context"
	database "github.com/alireza/identity/Database"
	"time"
	"fmt"
)




type TokenDb struct {
	Id   string  `db:"jit"`
	ExpiresAt time.Time `db:"expires_at"`
	AddedAt time.Time `db:"created_at"`
}
type TokenDataRepository struct {}

func NewTokenDataRepository() *TokenDataRepository {
	return &TokenDataRepository{
	}
}
func (t *TokenDataRepository) RevokeToken(ctx context.Context, jit string,CurrentDay string,expiresAt time.Time) error {
	db := database.GetDB()
	addedAt := time.Now().UTC()
	tableName := fmt.Sprintf("revoked_tokens_%s", CurrentDay)

	query :=fmt.Sprintf("INSERT INTO %s (jit, expires_at, created_at) VALUES (:jit, :expires_at, :created_at)", tableName)
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

func (t *TokenDataRepository) IsTokenRevoked(ctx context.Context, jit string ,tokenIssuanceDate string) (bool , error)  {
	db := database.GetDB()
	tableName := fmt.Sprintf("revoked_tokens_%s", tokenIssuanceDate)
	 query := fmt.Sprintf(`
        SELECT EXISTS(SELECT 1 FROM %s WHERE jit=?)
    `, tableName)

    var exists bool
    err := db.QueryRow(query, jit).Scan(&exists)
    return exists, err
}


