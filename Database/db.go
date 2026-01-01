package database

import (
	"log"
	"sync"

	 _ "github.com/go-sql-driver/mysql"

	"os"
	"github.com/jmoiron/sqlx"
)


var (
	instance *sqlx.DB
	ones sync.Once
)

func GetConfigurations()string{
	
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	return  user+":"+password+"@tcp("+host+":"+port+")/"+DBName+"?parseTime=true"
}

func GetDB() *sqlx.DB{
	ones.Do(func() {
		cfg := GetConfigurations()
		db , err := sqlx.Connect("mysql", cfg)
		if err!= nil{
			log.Fatal("Failed to connect to db: %v", err)

		}
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(0)
		instance = db
	})
	return  instance
}