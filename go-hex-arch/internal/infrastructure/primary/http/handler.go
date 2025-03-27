package http

import (
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/ports"

	"github.com/gofiber/fiber/v2"
)

type CalculatorHandler struct {
	service ports.CalculatorService
}

func NewCalculatorHandler(service ports.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: service}
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
