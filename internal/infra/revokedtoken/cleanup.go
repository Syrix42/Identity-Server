package revokedtoken

import (
	"context"
	"time"
	"github.com/alireza/identity/Database"
	"strings"
	"fmt"

)

func CleanupExpiredRevokedTokensTables(ctx context.Context , retentiondays int) error {
	db := database.GetDB()
	rows, err := db.QueryContext(ctx, `
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = DATABASE()
          AND table_name LIKE 'revoked_tokens_%'
    `)
    if err != nil {
        return err
    }
    defer rows.Close()

	cutoff := time.Now().AddDate(0, 0, -retentiondays)


	for rows.Next(){
		var tableName string
		if err:= rows.Scan(&tableName); err!= nil{
			continue
	}

	dataStr := strings.TrimPrefix (tableName , "revoked_tokens_")
	tableDate , err := time.Parse("20060102" , dataStr)
	if err != nil{
		continue
	}

	if tableDate.Before(cutoff){
		dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
		if _ , err := db.ExecContext (ctx , dropSQL); err != nil{
			fmt.Println("Failed to drop table" , tableName , err)
			continue
		}
		fmt.Println("Dropped table" , tableName)
	}
}
return nil
}

func StartDailyCleanup(ctx context.Context, retentionDays int) {
    go func() {
        for {
            now := time.Now().UTC()
            nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
            wait := time.Until(nextMidnight)

            time.Sleep(wait) 

            if err :=  CleanupExpiredRevokedTokensTables(ctx, retentionDays); err != nil {
                fmt.Println("Cleanup error:", err)
            }
        }
    }()
}
