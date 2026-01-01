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
	query := "INSERT INTO users (username , hashed_password , created_at , role, user_id) VALUES (:username,:hashed_password, :created_at, :role, :user_id)"
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
		//fmt.Println(err)
		return nil, err
		
	}
	return &domain.User{
		UserName: user.Username,
		UserID:   user.Id,
		Password: user.Password,
		Role:     user.Role,
	} , nil
	
}

func (m *MySQlDB) GetById(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{}, nil
}
