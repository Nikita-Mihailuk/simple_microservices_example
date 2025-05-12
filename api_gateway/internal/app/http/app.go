package http

import (
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/admin_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/user_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/delivery/http"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"time"
)

type App struct {
	router *fiber.App
	port   string
}

func NewApp(port string, userServiceClient *user_service.UserClient, adminServiceClient *admin_service.AdminClient) *App {
	router := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	router.Use(logger.New(), cors.New())

	handler := http.NewHandler(userServiceClient, adminServiceClient)
	handler.InitHandler(router)

	return &App{
		router: router,
		port:   port,
	}
}

func (a *App) Run() {
	if err := a.router.Listen(":" + a.port); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	if err := a.router.Shutdown(); err != nil {
		panic(err)
	}
}
