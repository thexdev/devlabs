package http

import (
	"scalt/internal/app/dto"
	"scalt/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type CalculatorHandler struct {
	service *service.CalculatorService
}

func NewCalculatorHandler(
	service service.CalculatorService,
) *CalculatorHandler {
	return &CalculatorHandler{service: &service}
}

func (h *CalculatorHandler) Calculate(c *fiber.Ctx) error {
	var input dto.CalculationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	output, err := h.service.PerformCalculation(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(output)
}
