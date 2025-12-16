package domain

type User struct {
	Name       string
	Password string
}

func NewUser(Name string , hash string) User{
	return  User{
		Name: Name,
		Password: hash,
	}
}