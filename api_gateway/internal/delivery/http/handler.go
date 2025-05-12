package http

import (
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/admin_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/user_service"
	v1 "github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/delivery/http/v1"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	userServiceClient  *user_service.UserClient
	adminServiceClient *admin_service.AdminClient
}

func NewHandler(userServiceClient *user_service.UserClient, adminServiceClient *admin_service.AdminClient) Handler {
	return Handler{
		userServiceClient:  userServiceClient,
		adminServiceClient: adminServiceClient,
	}
}

func (h *Handler) InitHandler(router fiber.Router) {
	router.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	api := router.Group("/api")

	handlerV1 := v1.NewHandlerV1(h.userServiceClient, h.adminServiceClient)
	handlerV1.InitHandlerV1(api)
}
