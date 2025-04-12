package usecases_test

import (
	"gohexarch/internal/app/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyUseCase_(t *testing.T) {
	uc := new(usecases.MultiplyUseCase)

	t.Run("multiply two possitive numbers", func(t *testing.T) {
		expected := float32(30)

		result, err := uc.Execute(2, 15)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("multiply with negative number", func(t *testing.T) {
		expected := float32(-30)

		result, err := uc.Execute(15, -2)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
