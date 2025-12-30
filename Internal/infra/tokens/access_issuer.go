package tokens

import (
	"context"
	"crypto/rsa"

	"github.com/alireza/identity/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)






type accessTokenClaims struct{
    Role string `json:"role"`
    jwt.RegisteredClaims
}


func  IssueAccessToken(ctx context.Context, token domain.Token , privateKey *rsa.PrivateKey) (string, error) {
    claims := accessTokenClaims{
        Role: token.Role,
        RegisteredClaims: jwt.RegisteredClaims{
			Subject: token.Subject ,
            Issuer: "aAuth",
            Audience: []string{"my-api"},
            IssuedAt: jwt.NewNumericDate(token.IssuedAt),
            ExpiresAt: jwt.NewNumericDate(token.ExpiresAt),
			ID: token.ID, 
        },
    }
    accesstoken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    return accesstoken.SignedString(privateKey)
}



