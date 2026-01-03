package tokens

import (
	"context"
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidRecoveryToken = errors.New("invalid recovery token")

func ValidateRecoveryToken(ctx context.Context, tokenstring string,  publickey *rsa.PublicKey) (*jwt.RegisteredClaims , error) {
	token, err := jwt.ParseWithClaims(tokenstring, &jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {

			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, ErrInvalidRecoveryToken
			}
			return publickey, nil

		})
	if err != nil || !token.Valid {
		return nil, ErrInvalidRecoveryToken
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrInvalidRecoveryToken
	}
	return claims, nil
}
