package domain


type UserRepo interface{
	 Save(u User) error
	 GetByName(u User) (*User,error)
}