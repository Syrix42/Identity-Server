package domain

import "errors"

var ErrCanNotAuthenticate = errors.New("To many Active Sessions")

type User struct {
	UserID        string
	UserName      string
	Password      string
	Role          string
	ActiveSession int
}

func NewUser(UserID string, Name string, password string, Role string) User {
	return User{
		UserID:        UserID,
		UserName:      Name,
		Password:      password,
		Role:          Role,
		ActiveSession: 1,
	}
}

func (u *User) ID() string {
	return u.UserID
}

func (u *User) Username() string {
	return u.UserName
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) CanAuthenticate() error {
	if u.ActiveSession >= 5 {
		return ErrCanNotAuthenticate
	}
	return nil
}
