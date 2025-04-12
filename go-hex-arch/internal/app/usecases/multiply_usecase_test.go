package usecases_test

import (
	"gohexarch/internal/app/usecases"
	mock_ "gohexarch/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMultiplyUseCase(t *testing.T) {
	mockRepo := new(mock_.MockRepository)

	uc := usecases.NewMultiplyUseCase(mockRepo)

	t.Run("multiply two possitive numbers", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(30)

		result, err := uc.Execute(2, 15)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("multiply with negative number", func(t *testing.T) {
		t.Cleanup(func() {
			mockRepo.Calls = nil
			mockRepo.ExpectedCalls = nil
		})

		mockRepo.On("Save", mock.Anything).Return(true, nil)

		expected := float32(-30)

		result, err := uc.Execute(15, -2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertNumberOfCalls(t, "Save", 1)
	})
}
