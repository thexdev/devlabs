package models_test

import (
	"gohexarch/internal/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockExpression(a float32, operator string, b float32) interface{} {
	expression := struct {
		A        float32
		B        float32
		Operator string
	}{
		A:        a,
		B:        b,
		Operator: operator,
	}

	return expression
}

func TestHistory(t *testing.T) {
	t.Run("initialized without error", func(t *testing.T) {
		history := models.NewHistory()

		assert.IsType(t, models.History{}, *history)
		assert.Len(t, history.Records, 0)
	})
}

func TestHistory_Push(t *testing.T) {
	t.Run("add operation history into records", func(t *testing.T) {
		history := models.NewHistory()

		mockExpression := mockExpression(1, "+", 2)

		result := float32(3)

		ok := history.Push(mockExpression, result)

		assert.True(t, ok)
		assert.Len(t, history.Records, 1)
	})
}

func TestHistory_All(t *testing.T) {
	t.Run("retrieve all record histories", func(t *testing.T) {
		history := models.NewHistory()

		mockExpression := mockExpression(1, "-", 1)

		result := float32(0)

		ok := history.Push(mockExpression, result)

		assert.True(t, ok)
		assert.Len(t, history.Records, 1)

		histories := history.All()

		assert.Len(t, histories, 1)
	})
}
