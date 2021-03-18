package internal

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/lucabecci/parking-lot/pkg/models"
	"github.com/lucabecci/parking-lot/services/auth/service"
)

type logMW struct {
	logger log.Logger
	service.Service
}

func NewLoggingMiddleware(logger log.Logger, next service.Service) logMW {
	return logMW{logger, next}
}

func (mw logMW) Register(ctx context.Context, email, password, confirmPassword string) (usr models.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "register",
			"input", email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	usr, err = mw.Service.Register(ctx, email, password, confirmPassword)
	return
}

func (mw logMW) Login(ctx context.Context, email, password string) (token string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "login",
			"input", email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	token, err = mw.Service.Login(ctx, email, password)
	return
}
