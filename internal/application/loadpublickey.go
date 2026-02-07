package application

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPublicKey() (*rsa.PublicKey, error) {
	
	keyPath := os.Getenv("PUBLIC_KEY_PATH")
	if keyPath == "" {
		return nil, errors.New("PUBLIC_KEY_PATH is not set")
	}


	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}


	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
