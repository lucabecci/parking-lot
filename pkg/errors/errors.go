package errors

import "errors"

var (
	ErrInvalidPassword = errors.New("Invalid Password")
	ErrEmailInUse      = errors.New("Email already in use")
	ErrInvalidToken    = errors.New("Token invalid")
)
