package http

import (
	"gohexarch/internal/app/ports/primary"
	"gohexarch/internal/app/ports/secondary"

	"github.com/gofiber/fiber/v2"
)

type HistoryResponse struct {
	ID         string            `json:"id"`
	Expression secondary.Expression `json:"expression"`
	Result     float32              `json:"result"`
}

type HistoryHandler struct {
	retrieve primary.History
}

func NewHistoryHandler(retrieve primary.History) *HistoryHandler {
	return &HistoryHandler{
		retrieve: retrieve,
	}
}

func (h *HistoryHandler) All(c *fiber.Ctx) error {
	result, err := h.retrieve.Execute()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "histories not found",
		})
	}

	histories := []HistoryResponse{}
	for _, history := range result {
		histories = append(histories, HistoryResponse{
			ID: history.ID,
			Expression: history.Expression,
			Result: history.Result,
		})
	}

	return c.JSON(histories)
}
