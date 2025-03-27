package services

import (
	"context"
	"errors"
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/ports"
)

type CalculatorServiceImpl struct{}

func NewCalculatorService() ports.CalculatorService {
	return &CalculatorServiceImpl{}
}

func (s *CalculatorServiceImpl) Calculate(ctx context.Context, req models.CalculationRequest) (models.CalculationResponse, error) {
	switch req.Operator {
	case "+":
		return models.CalculationResponse{Result: req.A + req.B}, nil
	case "-":
		return models.CalculationResponse{Result: req.A - req.B}, nil
	case "*":
		return models.CalculationResponse{Result: req.A * req.B}, nil
	case "/":
		if req.B == 0 {
			return models.CalculationResponse{Error: "division by zero"}, errors.New("division by zero")
		}
		return models.CalculationResponse{Result: req.A / req.B}, nil
	default:
		return models.CalculationResponse{Error: "invalid operator"}, errors.New("invalid operator")
	}
}
