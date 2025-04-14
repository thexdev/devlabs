package http

import (
	"gohexarch/internal/app/ports/primary"

	"github.com/gofiber/fiber/v2"
)

type CalculatorRequest struct {
	A        float32 `json:"a"`
	B        float32 `json:"b"`
	Operator string  `json:"operator"`
}

type CalculatorResponse struct {
	Expression CalculatorRequest `json:"expression"`
	Result     float32           `json:"result"`
	Error      string            `json:"error,omitempty"`
}

type CalculatorHandler struct {
	Add      primary.Add
	Subtract primary.Subtract
	Devide   primary.Devide
	Multiply primary.Multiply
}

func NewCalculatorHandler(
	add primary.Add,
	subtract primary.Subtract,
	devide primary.Devide,
	multiply primary.Multiply,
) *CalculatorHandler {
	return &CalculatorHandler{
		Add:      add,
		Subtract: subtract,
		Devide:   devide,
		Multiply: multiply,
	}
}

func (h *CalculatorHandler) Compute(c *fiber.Ctx) error {
	var (
		req = new(CalculatorRequest)
		res = new(CalculatorResponse)
	)

	if err := c.BodyParser(&req); err != nil {
		res.Expression = *req
		res.Error = "invalid expression"

		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	switch req.Operator {
	case "+":
		result, _ := h.Add.Execute(req.A, req.B)

		res.Expression = *req
		res.Result = result
	case "-":
		result, _ := h.Subtract.Execute(req.A, req.B)

		res.Expression = *req
		res.Result = result
	case "*":
		result, _ := h.Multiply.Execute(req.A, req.B)

		res.Expression = *req
		res.Result = result
	case "/":
		result, err := h.Devide.Execute(req.A, req.B)

		if err != nil {
			res.Error = err.Error()
		}

		res.Expression = *req
		res.Result = result
	default:
		res.Expression = *req
		res.Error = "invalid operator"
	}

	if res.Error != "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(res)
	}

	return c.JSON(res)
}
