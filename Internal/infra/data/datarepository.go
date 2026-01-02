package data

import (
	"context"
	"fmt"
	"time"

	database "github.com/alireza/identity/Database"
	"github.com/alireza/identity/internal/domain"
)

type UserDb struct {
	Id        string    `db:"id"`
	UserId        string    `db:"user_id"`
	Username  string    `db:"username"`
	Password  string    `db:"hashed_password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	ActiveUsers int       `db:"active_sessions"`
}

type MySQlDB struct{}

func NewMYSQLDb() *MySQlDB {
	return &MySQlDB{}
}

func (m *MySQlDB) Save(ctx context.Context, u domain.User) error {
	db := database.GetDB()

	user := UserDb{
		UserId:        u.ID(),
		Username:  u.UserName,
		Password:  u.Password,
		CreatedAt: time.Now().UTC(),
		Role:      u.Role,
	}
	query := "INSERT INTO users (username , hashed_password , created_at , role, user_id , active_sessions) VALUES (:username,:hashed_password, :created_at, :role, :user_id, :active_sessions)"
	_, err := db.NamedExec(query, user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *MySQlDB) GetByName(ctx context.Context, name string) (*domain.User, error) {
	user := UserDb{}
	db := database.GetDB()
	query := "SELECT * FROM users WHERE username = ?"
	err := db.Get(&user, query, name)
	if err != nil {
	
		return nil, err
		
	}
	return &domain.User{
		UserName: user.Username,
		UserID:   user.UserId,
		Password: user.Password,
		Role:     user.Role,
		ActiveSession: user.ActiveUsers,
	} , nil
	
}

func (m *MySQlDB) GetById(ctx context.Context, id string) (*domain.User, error) {
	user := UserDb{}
	db := database.GetDB()
	query := "SELECT * FROM users WHERE user_id = ?"
	err := db.Get(&user , query , id)
	if err!=nil{
		return nil , err
	}
	return &domain.User{
		UserName: user.Username,
		Password: user.Password,
		UserID: user.UserId,
		Role: user.Role,
		ActiveSession: user.ActiveUsers,
	} , nil

}

func (m *MySQlDB) IncrementActiveSessions(userID string) error {
	fmt.Println(userID)
	db := database.GetDB()
	query := "UPDATE users SET active_sessions = active_sessions + 1 WHERE user_id = ?"
	_, err := db.Exec(query, userID)

	if err!= nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (m *MySQlDB) DecrementActiveSessions(userID string) error {
	
	
	db := database.GetDB()
	query := "UPDATE users SET active_sessions = active_sessions - 1 WHERE user_id = ?"
	_, err := db.Exec(query, userID)
	if err!=nil{
		
		return err
	}
	res, err := db.Exec(query, userID)
    rows, _ := res.RowsAffected()
    fmt.Println("Rows affected:", rows)
	return nil 
}
