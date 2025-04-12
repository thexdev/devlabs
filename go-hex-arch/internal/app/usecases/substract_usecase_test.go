package usecases_test

import (
	"gohexarch/internal/app/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstractUseCase(t *testing.T) {
	uc := new(usecases.SubstractUseCase)

	t.Run("substract two possitive numbers", func(t *testing.T) {
		expected := float32(20)

		result, err := uc.Execute(50, 30)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("subtract with negative number", func(t *testing.T) {
		expected := float32(5)

		result, err := uc.Execute(0, -5)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
