package service

import (
	"context"

	"github.com/lucabecci/parking-lot/pkg"
	"github.com/lucabecci/parking-lot/pkg/models"
	"github.com/lucabecci/parking-lot/pkg/repository"
	"github.com/lucabecci/parking-lot/pkg/security"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password, confirmPassword string) (models.User, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

type service struct {
	repository repository.UserRepository
}

func GetService(repository repository.UserRepository) *service {
	return &service{repository: repository}
}

func (s *service) Register(ctx context.Context, email, password, confirmPassword string) (models.User, error) {
	if password != confirmPassword {
		return models.User{}, pkg.ErrPasswordNotEqual
	}
	usr, err := s.repository.Create(email, password)
	if err != nil {
		return models.User{}, pkg.ErrToCreate
	}
	return usr, nil
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	match := user.PasswordMatch(password)
	if match == false {
		return "", pkg.ErrInvalidPassword
	}
	token, err := security.NewToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", pkg.ErrInvalidToken
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", pkg.ErrInvalidToken
	}
	return tData["id"].(string), nil
}
