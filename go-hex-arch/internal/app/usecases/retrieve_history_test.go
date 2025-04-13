package usecases_test

import (
	"gohexarch/internal/app/ports/secondary"
	"gohexarch/internal/app/usecases"
	mock_ "gohexarch/internal/test/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveHistory(t *testing.T) {
	t.Run("initialize struct without error", func(t *testing.T) {
		mockRepo := new(mock_.MockRepository)

		uc := usecases.NewRetrieveHistoryUseCase(mockRepo)

		assert.IsType(t, usecases.RetrieveHistoryUseCase{}, *uc)
	})
}

func TestRetrieveHistory_Execute(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewRetrieveHistoryUseCase(mockRepo)

	t.Run("returns empty slice when theres no record", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("All").Return([]secondary.Record{})

		result, err := uc.Execute()

		assert.NoError(t, err)
		assert.Len(t, result, 0)

		mockRepo.AssertNumberOfCalls(t, "All", 1)
	})

	t.Run("returns slice of calculation histories", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockExpression := secondary.Expression{
			A:        1,
			B:        2,
			Operator: "+",
		}

		mockRecord := secondary.Record{
			ID:         uuid.NewString(),
			Expression: mockExpression,
			Result:     float32(3),
		}

		mockHistories := []secondary.Record{mockRecord}

		mockRepo.On("All").Return(mockHistories)

		result, err := uc.Execute()

		assert.NoError(t, err)
		assert.Len(t, result, 1)

		mockRepo.AssertNumberOfCalls(t, "All", 1)
	})
}
