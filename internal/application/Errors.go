package application

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("User Aldready Exists")
var ErrCanNotAuthenticate = errors.New("To many Active Sessions")
var ErrInvalidToken = errors.New("Invalid Token")
var ErrTokenAlreadyRevoked = errors.New("Token Already Revoked")