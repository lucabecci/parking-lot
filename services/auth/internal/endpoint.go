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
