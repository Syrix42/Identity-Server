package domain



type TokenRepository interface{
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
}