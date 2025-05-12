package v1

import (
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/admin_service"
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/clients/user_service"
	"github.com/gofiber/fiber/v3"
)

type HandlerV1 struct {
	userServiceClient  *user_service.UserClient
	adminServiceClient *admin_service.AdminClient
}

func NewHandlerV1(userServiceClient *user_service.UserClient, adminServiceClient *admin_service.AdminClient) *HandlerV1 {
	return &HandlerV1{
		userServiceClient:  userServiceClient,
		adminServiceClient: adminServiceClient,
	}
}

func (h *HandlerV1) InitHandlerV1(api fiber.Router) {
	v1 := api.Group("/v1")
	h.RegisterUserRouts(v1)
	h.RegisterAdminRouts(v1)
}
