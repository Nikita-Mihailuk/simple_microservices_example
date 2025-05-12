package v1

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (h *HandlerV1) RegisterAdminRouts(v1 fiber.Router) {
	adminGroup := v1.Group("/admins")

	adminGroup.Delete("/users/:id", h.deleteUser)
}

func (h *HandlerV1) deleteUser(c fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	isDeleted, err := h.adminServiceClient.DeleteUserById(c.Context(), int64(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"isDeleted": isDeleted,
	})
}
