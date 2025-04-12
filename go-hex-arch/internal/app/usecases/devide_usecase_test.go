package usecases_test

import (
	"gohexarch/internal/app/usecases"
	mock_ "gohexarch/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDevideUseCase(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewDevideUseCase(mockRepo)

	t.Run("devide by zero", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(0)

		result, err := uc.Execute(1, 0)

		assert.Error(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("devide two possitive numbers", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(5)

		result, err := uc.Execute(10, 2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("devide with negative number", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(-5)

		result, err := uc.Execute(10, -2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})
}
