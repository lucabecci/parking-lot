package internal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/lucabecci/parking-lot/pkg"
	"github.com/lucabecci/parking-lot/services/auth/service"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(svc service.Service, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	registerHandler := httptransport.NewServer(
		MakeRegisterEndpoint(svc),
		decodeRegisterRequest,
		pkg.EncodeResponse,
		options...,
	)

	loginHandler := httptransport.NewServer(
		MakeLoginEndpoint(svc),
		decodeLoginRequest,
		pkg.EncodeResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/login").Handler(loginHandler)
	r.Methods("POST").Path("/register").Handler(registerHandler)
	return r
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeError with nil error")
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case pkg.ErrEmailNotExists:
		return http.StatusNotFound
	case pkg.ErrInvalidToken:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
