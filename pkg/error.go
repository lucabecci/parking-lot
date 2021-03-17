package pkg

import "errors"

var (
	ErrInvalidPassword  = errors.New("Invalid Password")
	ErrEmailInUse       = errors.New("Email already in use")
	ErrEmailNotExists   = errors.New("Email not registered")
	ErrInvalidToken     = errors.New("Token invalid")
	ErrToCreate         = errors.New("Error to create your user")
	ErrPasswordNotEqual = errors.New("Password not equal")
)
