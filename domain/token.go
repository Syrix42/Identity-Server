package domain

import "time"

type Type string

const (
	Access Type = "access"
	Recovery Type = "recovery"

)


type Token struct{
	ID string
	Subject string
	TokenType Type
	IssuedAt  time.Time
	ExpiresAt time.Time
	Role string
}

func NewAcessToken(Id string , Subject string , Ttype Type , iss time.Time ,Role string) *Token{
	return &Token{
		ID: Id,
		Subject: Subject,
		TokenType: Ttype,
		IssuedAt: iss,
		ExpiresAt: iss.Add(15 * time.Minute),
		Role: Role,
	}
}

func NewRecoveryToken(Id string , Subject string , Ttype Type , iss time.Time ) *Token{
	return &Token{
		ID: Id,
		Subject: Subject,
		TokenType: Ttype,
		IssuedAt: iss,
		ExpiresAt: iss.Add(time.Hour * 24 * 7),
	}
}