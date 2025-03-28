package service_test

import (
	"scalt/internal/app/dto"
	"scalt/internal/app/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorService_PerformCalculation(t *testing.T) {
	calculatorService := service.NewCalculatorService()

	testCases := []struct{
		name string
		input dto.CalculationInput
		expected float64
	} {
		{
			name: "valid calculation for addtion",
			input: dto.CalculationInput{
				LeftOperand: 10,
				Operator: "+",
				RightOperand: 5,
			},
			expected: 15.0,
		},
		{
			name: "valid calculation for subtraction",
			input: dto.CalculationInput{
				LeftOperand: 10,
				Operator: "-",
				RightOperand: 5,
			},
			expected: 5.0,
		},
		{
			name: "valid calculation for multiplication",
			input: dto.CalculationInput{
				LeftOperand: 10,
				Operator: "*",
				RightOperand: 5,
			},
			expected: 50.0,
		},
		{
			name: "valid calculation for division",
			input: dto.CalculationInput{
				LeftOperand: 10,
				Operator: "/",
				RightOperand: 5,
			},
			expected: 2.0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			input := dto.CalculationInput{
				LeftOperand: test.input.LeftOperand,
				Operator: test.input.Operator,
				RightOperand: test.input.RightOperand,
			}

			output, err := calculatorService.PerformCalculation(input)

			assert.NoError(t, err)
			assert.Equal(t, test.expected, output.Result)
		})
	}
}
