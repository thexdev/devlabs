package ports

import (
	"context"
	"gohexarch/internal/domain/models"
)

// PRIMARY PORT
type CalculatorService interface {
	Calculate(ctx context.Context, req models.CalculationRequest) (models.CalculationResponse, error)
}
