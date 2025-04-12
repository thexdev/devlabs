package models_test

import (
	"gohexarch/internal/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	calculator := models.NewCalculator()

	t.Run("addition", func(t *testing.T) {
		t.Run("add two possitive numbers", func(t *testing.T) {
			expected := float32(3)

			result, err := calculator.Add(1, 2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})

		t.Run("add with negative number", func(t *testing.T) {
			expected := float32(3)

			result, err := calculator.Add(5, -2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("subtraction", func(t *testing.T) {
		t.Run("subtract two possitive number", func(t *testing.T) {
			expected := float32(5)

			result, err := calculator.Subtract(10, 5)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})

		t.Run("subtract with negative number", func(t *testing.T) {
			expected := float32(10)

			result, err := calculator.Subtract(5, -5)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("multiplication", func(t *testing.T) {
		t.Run("multiply two possitive numbers", func(t *testing.T) {
			expected := float32(20)

			result, err := calculator.Multiply(10, 2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})

		t.Run("multiply with negative numbers", func(t *testing.T) {
			expected := float32(-20)

			result, err := calculator.Multiply(10, -2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("division", func(t *testing.T) {
		t.Run("cannot divide by zero", func(t *testing.T) {
			expected := float32(0)

			result, err := calculator.Devide(1, 0)

			assert.Error(t, err)
			assert.Equal(t, expected, result)
		})

		t.Run("devide two possitive numbers", func(t *testing.T) {
			expected := float32(25)

			result, err := calculator.Devide(50, 2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})

		t.Run("devide with negative numbers", func(t *testing.T) {
			expected := float32(-25)

			result, err := calculator.Devide(50, -2)

			assert.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	})
}
