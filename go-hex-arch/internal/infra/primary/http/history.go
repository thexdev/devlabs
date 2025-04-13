package http

import (
	"gohexarch/internal/app/ports/primary"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	retrieve primary.History
}

func NewHistoryHandler(retrieve primary.History) *HistoryHandler {
	return &HistoryHandler{
		retrieve: retrieve,
	}
}

func (h *HistoryHandler) All(c *fiber.Ctx) error {
	histories, err := h.retrieve.Execute()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "histories not found",
		})
	}

	return c.JSON(histories)
}
