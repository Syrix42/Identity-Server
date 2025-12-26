package tokens

import (
	"context"
	"crypto/rsa"

	"github.com/alireza/identity/domain"
	"github.com/golang-jwt/jwt/v5"
)



func IssueRecoveryToken(ctx context.Context,Token domain.Token ,  privateKey *rsa.PrivateKey) (string , error){
	claims := jwt.RegisteredClaims{
		Subject: Token.Subject,
		ID: Token.ID,
		ExpiresAt: jwt.NewNumericDate(Token.ExpiresAt),
		IssuedAt: jwt.NewNumericDate(Token.IssuedAt),
		Issuer: "aAUT",
		Audience: []string{"AauthService"},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	
    return  jwtToken.SignedString(privateKey)

	
}