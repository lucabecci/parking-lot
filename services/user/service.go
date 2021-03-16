package user

import (
	"context"
	"errors"
)

type Service interface {
	ValidateUser(ctx context.Context, email string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("User Invalid")
	ErrInvalidToken = errors.New("Token Invalid")
)

type service struct{}

func NewService() *service {
	return &service{}
}

// func (s *service) ValidateUser(ctx context.Context, email, password string) (string, error) {

// }
