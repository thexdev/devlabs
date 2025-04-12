package usecases_test

import (
	"gohexarch/internal/app/usecases"
	mock_ "gohexarch/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubstractUseCase(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewSubtractUseCase(mockRepo)

	t.Run("substract two possitive numbers", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(20)

		result, err := uc.Execute(50, 30)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("subtract with negative number", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(5)

		result, err := uc.Execute(0, -5)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})
}
