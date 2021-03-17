package pkg

import "errors"

var (
	//Email
	ErrEmailInUse     = errors.New("Email already in use")
	ErrEmailNotExists = errors.New("Email not registered")
	//Creation
	ErrToCreate = errors.New("Error to create your user")
	//Password
	ErrPasswordNotEqual = errors.New("Password not equal")
	ErrInvalidPassword  = errors.New("Invalid Password")
	//Token
	ErrInvalidToken  = errors.New("Token invalid")
	ErrToCreateToken = errors.New("Error to create your token")
)
