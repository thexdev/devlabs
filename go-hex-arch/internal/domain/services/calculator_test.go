package services_test

import (
	"context"
	"gohexarch/internal/domain/models"
	"gohexarch/internal/domain/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorService(t *testing.T) {
	// Initialize the domain service (no mocks needed)
	service := services.NewCalculatorService()

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