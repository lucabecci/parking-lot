package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lucabecci/parking-lot/pkg/repository"
	"github.com/lucabecci/parking-lot/services/auth/internal"
	"github.com/lucabecci/parking-lot/services/auth/service"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	uri := os.Getenv("DB_URI")
	data, err := internal.ConnectDB(uri)
	if err != nil {
		fmt.Println("Error in database module")
	}
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)

	svc := internal.NewLoggingMiddleware(
		logger,
		service.GetService(repository.UserRepository{Database: data.Db}),
	)
	r := internal.NewHTTPServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}
