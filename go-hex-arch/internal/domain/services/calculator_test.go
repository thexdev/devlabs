package services_test

import (
	"context"
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/services"
	"gohexarch/internal/infrastructure/secondary/inmemory"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the core calculator logic AND history recording
func TestCalculatorService_WithHistory(t *testing.T) {
	// Setup: Create in-memory repo and service
	historyRepo := inmemory.NewHistoryRepository()
	service := services.NewCalculatorService(historyRepo)

	// Define test cases
	tests := []struct {
		name string
		request models.CalculationRequest
		expected float64
		hasError bool
	}{
		{
			name: "Addition",
			request: models.CalculationRequest{A: 10, B: 5, Operator: "+"},
			expected: 15.0,
		},
		{
			name: "Division by zero",
			request: models.CalculationRequest{A: 10, B: 0, Operator: "/"},
			expected: 0,
			hasError: true,
		},
		{
			name: "Invalid operator",
			request: models.CalculationRequest{A: 10, B: 5, Operator: "?"},
			expected: 0,
			hasError: true,
		},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := service.Calculate(context.Background(), tc.request)

			if tc.hasError {
				assert.Error(t, err)
				assert.NotEmpty(t, resp.Error)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, resp.Result)
			}
		})
	}
}

func TestCalculatorService_History(t *testing.T) {
	// Mock repository
	mockRepo := new(MockHistoryRepository)
	service := services.NewCalculatorService(mockRepo)

	// Test calculation + history save
	req := models.CalculationRequest{A: 10, B: 5, Operator: "+"}
	resp, err := service.Calculate(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, 15.0, resp.Result)
	assert.True(t, mockRepo.AddEntryCalled) // Verify AddEntry was called
}


// Mock Repository
type MockHistoryRepository struct {
	AddEntryCalled bool
}

func (m *MockHistoryRepository) AddEntry(
	ctx context.Context,
	entry models.CalculationHistory,
) error {
	m.AddEntryCalled = true
	return nil
}

func (m *MockHistoryRepository) GetHistory(
	ctx context.Context,
) ([]models.CalculationHistory, error) {
	return nil, nil
}

func TestCalculatorService_ConcurrencyHistory(t *testing.T) {
	repo := inmemory.NewHistoryRepository()
	service := services.NewCalculatorService(repo)

	// Run 10 concurrent calculations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			req := models.CalculationRequest{A: float64(i), B: 1, Operator: "+"}
			_, _ = service.Calculate(context.Background(), req)
		}(i)
	}
	wg.Wait()

	// Verify all entries were recorded
	history, err := repo.GetHistory(context.Background())
	assert.NoError(t, err)
	assert.Len(t, history, 10)
}
