package domain


type UserRepo interface{
	 Save(u User) error
	 GetByName(name string) (*User,error)
}