package usecases_test

import (
	"gohexarch/internal/app/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDevideUseCase(t *testing.T) {
	uc := new(usecases.DevideUseCase)

	t.Run("devide by zero", func(t *testing.T) {
		expected := float32(0)

		result, err := uc.Execute(1, 0)

		assert.Error(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("devide two possitive numbers", func(t *testing.T) {
		expected := float32(5)

		result, err := uc.Execute(10, 2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("devide with negative number", func(t *testing.T) {
		expected := float32(-5)

		result, err := uc.Execute(10, -2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
