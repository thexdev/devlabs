package ports

import (
	"context"
	"gohexarch/internal/domain/models"
)

type HistoryRepository interface {
	AddEntry(ctx context.Context, entry models.CalculationHistory) error
	GetHistory(ctx context.Context) ([]models.CalculationHistory, error)
}
