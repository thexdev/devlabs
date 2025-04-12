package usecases_test

import (
	"gohexarch/internal/app/usecases"
	mock_ "gohexarch/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUseCase_CalculateTwoPossitiveNumbers(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewAddUseCase(mockRepo)

	t.Run("add two possitive numbers", func(t *testing.T) {
		mockRepo.On("Save", mock.Anything).Return(true, nil)
		
		expected := float32(10)

		result, err := uc.Execute(7, 3)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)

		assert.NoError(t, err)
		assert.Equal(t, result, expected)
	})
}

func TestAddUseCase_CalculateWithNegativeNumber(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewAddUseCase(mockRepo)

	t.Run("add with negative number", func(t *testing.T) {
		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(5)

		result, err := uc.Execute(7, -2)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)

		assert.NoError(t, err)
		assert.Equal(t, result, expected)
	})
}
