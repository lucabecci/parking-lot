package service

import "context"

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password, confirmPassword string) error
	ValidateToken(ctx context.Context, token string) error
}
