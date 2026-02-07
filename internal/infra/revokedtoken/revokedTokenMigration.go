package revokedtoken

import(
	"sync"
	"fmt"
	"time"
	"context"
	"github.com/alireza/identity/Database"
)

var (
 
    activeDateMutex sync.RWMutex
    currentActiveDate string 
)


func EnsureRevokedTokenTable(ctx context.Context, tokenIssuanceDate time.Time) (string , error){
	tokenDay := tokenIssuanceDate.Format("20060102")
	activeDateMutex.Lock()
	if currentActiveDate == "" {
		currentActiveDate = tokenDay
	}
	defer activeDateMutex.Unlock()
	if tokenDay >= currentActiveDate {
		currentActiveDate = tokenDay
	
	tableName := fmt.Sprintf("revoked_tokens_%s", tokenDay)
	 createSQL := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id INT AUTO_INCREMENT PRIMARY KEY,
            jit VARCHAR(512) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            expires_at TIMESTAMP NOT NULL,
            UNIQUE INDEX idx_jit (jit),
            INDEX idx_expires_at (expires_at)
        );`, tableName)
	db := database.GetDB()
	_, err := db.Exec(createSQL)
	if err != nil {
		return "", err
	}
	return tokenDay , nil
}
return  currentActiveDate , nil
}

