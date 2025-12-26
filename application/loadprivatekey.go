package application

import (
	"crypto/rsa"
	"os"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func PathLoader() string{
	  keyPath  := os.Getenv("PRIVATE_KEY_PATH")
	  return  keyPath
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error){
	keybytes , err := os.ReadFile(path)
	if err!= nil{
	    return nil, fmt.Errorf("failed to read private key from %s: %w", path, err)

	}

	key , err := jwt.ParseRSAPrivateKeyFromPEM(keybytes)
	if err != nil{
		return  nil , err
	}
	return  key , nil

}