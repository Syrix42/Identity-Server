package application



type PasswordHasher interface{
	Hash(plaintext string) (string ,error)
}