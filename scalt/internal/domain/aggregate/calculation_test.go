package aggregate_test

import (
	"scalt/internal/domain/aggregate"
	"scalt/internal/domain/entity"
	"scalt/internal/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAggregate(t *testing.T) {
	testCases := []struct {
		name         string
		leftOperand  float64
		operator     string
		rightOperand float64
		expected     float64
	}{
		{
			name: "perform addtion",
			leftOperand: 10,
			operator: "+",
			rightOperand: 5,
			expected: 15.0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			leftOperand := valueobject.NewOperand(test.leftOperand)
			rightOperand := valueobject.NewOperand(test.rightOperand)
			operator, _ := valueobject.NewOperator(test.operator)
			operation := entity.NewOperation(
				leftOperand,
				operator,
				rightOperand,
			)
			calculation := aggregate.NewCalculation(*operation)

			calculation.Compute()

			assert.Equal(t, calculation.Result(), 15.0)
		})
	}
}
