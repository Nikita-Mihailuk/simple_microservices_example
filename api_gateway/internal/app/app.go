package app

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/app/http"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/admin_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/user_service"
	"time"
)

type App struct {
	HTTPServer *http.App
}

func NewApp(port string,
	userServiceTimeout time.Duration,
	userServiceAddr string,
	adminServiceTimeout time.Duration,
	adminServiceAddr string,
) *App {

	ctxUser, cancelUser := context.WithTimeout(context.Background(), userServiceTimeout)
	defer cancelUser()
	userServiceClient, err := user_service.NewUserClient(ctxUser, userServiceAddr)
	if err != nil {
		panic(err)
	}

	ctxAdmin, cancelAdmin := context.WithTimeout(context.Background(), adminServiceTimeout)
	defer cancelAdmin()
	adminServiceClient, err := admin_service.NewAdminClient(ctxAdmin, adminServiceAddr)
	if err != nil {
		panic(err)
	}

	httpApp := http.NewApp(port, userServiceClient, adminServiceClient)
	return &App{
		HTTPServer: httpApp,
	}
}
