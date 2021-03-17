package service

import (
	"context"

	"github.com/lucabecci/parking-lot/pkg"
	"github.com/lucabecci/parking-lot/pkg/models"
	"github.com/lucabecci/parking-lot/pkg/repository"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password, confirmPassword string) (models.User, error)
	ValidateToken(ctx context.Context, token string) error
}

type service struct {
	respository repository.UserRepository
}

func GetService() *service {
	return &service{}
}

func (s *service) Register(ctx context.Context, email, password, confirmPassword string) (models.User, error) {
	if password != confirmPassword {
		return models.User{}, pkg.ErrPasswordNotEqual
	}
	usr, err := s.respository.Create(email, password)
	if err != nil {
		return models.User{}, pkg.ErrToCreate
	}
	return usr, nil
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.respository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	match := user.PasswordMatch(password)
	if match == false {
		return "", pkg.ErrInvalidPassword
	}
	return "", nil
}

func (s *service) ValidateToken(ctx context.Context, token string) error {
	return nil
}
