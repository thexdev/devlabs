package services

import (
	"context"
	"errors"
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/ports"
	"time"

	"github.com/google/uuid"
)

type CalculatorServiceImpl struct{
	historyRepo ports.HistoryRepository // Inject repository
}

func NewCalculatorService(r ports.HistoryRepository) ports.CalculatorService {
	return &CalculatorServiceImpl{historyRepo: r}
}

func (s *CalculatorServiceImpl) Calculate(ctx context.Context, req models.CalculationRequest) (models.CalculationResponse, error) {
	var resp models.CalculationResponse

	// Business logic
	switch req.Operator {
	case "+":
		resp.Result = req.A + req.B
	case "-":
		resp.Result = req.A - req.B
	case "*":
		resp.Result = req.A * req.B
	case "/":
		if req.B == 0 {
			resp.Error = "division by zero"
			return resp, errors.New("division by zero")
		}
		resp.Result = req.A / req.B
	default:
		resp.Error = "invalid operator"
		return resp, errors.New("invalid operator")
	}

	// Save to history
	entry := models.CalculationHistory{
		ID: uuid.NewString(),
		Request: req,
		Response: resp,
		Timestamp: time.Now().Unix(),
	}
	if err := s.historyRepo.AddEntry(ctx, entry); err != nil {
		return resp, err
	}

	return resp, nil
}
