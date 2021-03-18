package internal

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucabecci/parking-lot/pkg/models"
	"github.com/lucabecci/parking-lot/services/auth/service"
)

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RegisterResponse struct {
	User models.User `json:"user,omitempty"`
	Err  string      `json:"error,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"error,omitempty"`
}

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

func MakeRegisterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		usr, err := svc.Register(ctx, req.Email, req.Password, req.ConfirmPassword)
		if err != nil {
			return RegisterResponse{models.User{}, err.Error()}, err
		}
		return RegisterResponse{usr, ""}, nil
	}
}

func MakeLoginEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := svc.Login(ctx, req.Email, req.Password)
		if err != nil {
			return LoginResponse{"", err.Error()}, err
		}
		return LoginResponse{token, ""}, nil
	}
}

func MakeValidateTokenEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateTokenRequest)
		id, err := svc.ValidateToken(ctx, req.Token)
		if err != nil {
			return ValidateTokenResponse{"", err.Error()}, err
		}
		return ValidateTokenResponse{id, ""}, nil
	}
}
