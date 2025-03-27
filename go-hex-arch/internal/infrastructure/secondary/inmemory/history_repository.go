package inmemory

import (
	"context"
	"gohexarch/internal/domain/models"
	"sync"
)

type HistoryRepository struct {
	entries []models.CalculationHistory
	mu sync.Mutex // Thread-safe
}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{
		entries: make([]models.CalculationHistory, 0),
	}
}

func (r *HistoryRepository) AddEntry(
	ctx context.Context,
	entry models.CalculationHistory,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries = append(r.entries, entry)
	return nil
}

func (r *HistoryRepository) GetHistory(
	ctx context.Context,
) ([]models.CalculationHistory, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.entries, nil
}
