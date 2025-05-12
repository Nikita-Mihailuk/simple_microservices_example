package v1

import (
	"github.com/Nikita-Mihailuk/simple_microservices_example/api_gateway/internal/domain/dto"
	"github.com/gofiber/fiber/v3"
)

func (h *HandlerV1) RegisterUserRouts(v1 fiber.Router) {
	userGroup := v1.Group("/users")

	userGroup.Get("/", h.getUsers)
	userGroup.Post("/", h.createUser)
}

func (h *HandlerV1) getUsers(c fiber.Ctx) error {
	users, err := h.userServiceClient.GetUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	return c.Status(200).JSON(users)
}

func (h *HandlerV1) createUser(c fiber.Ctx) error {
	var input dto.InputUser
	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	userID, err := h.userServiceClient.CreateUser(c.Context(), input.Name, input.Age)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": userID,
	})
}
