package entity_test

import (
	"scalt/internal/domain/entity"
	"scalt/internal/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperation(t *testing.T) {
	t.Run("create valid operation", func(t *testing.T) {
		leftOperand := valueobject.NewOperand(5)
		rightOperand := valueobject.NewOperand(3)
		operator, _ := valueobject.NewOperator("+")
		operation := entity.NewOperation(leftOperand, operator, rightOperand)

		assert.Equal(t, operation.LeftOperand.Value(), 5.0)
		assert.Equal(t, operation.RightOperand.Value(), 3.0)
		assert.Equal(t, operator.Symbol(), "+")
	})
}
