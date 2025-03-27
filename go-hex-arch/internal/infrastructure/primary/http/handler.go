package http

import (
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/ports"

	"github.com/gofiber/fiber/v2"
)

type CalculatorHandler struct {
	service ports.CalculatorService
	historyRepo ports.HistoryRepository
}

func NewCalculatorHandler(
	service ports.CalculatorService,
	repo ports.HistoryRepository,
) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
		historyRepo: repo,
	}
}

func (h CalculatorHandler) Calculate(c *fiber.Ctx) error {
	var req models.CalculationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CalculationResponse{
			Error: "invalid request",
		})
	}

	// Call domain service
	resp, err := h.service.Calculate(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	return c.JSON(resp)
}

func (h *CalculatorHandler) GetHistory(c *fiber.Ctx) error {
	histories, err := h.historyRepo.GetHistory(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch history",
		})
	}

	if len(histories) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "histories not found",
		})
	}

	return c.JSON(histories)
}
