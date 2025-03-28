package service

import (
	"scalt/internal/app/dto"
	"scalt/internal/domain/aggregate"
	"scalt/internal/domain/entity"
	"scalt/internal/domain/valueobject"
)

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) PerformCalculation(
	input dto.CalculationInput,
) (dto.CalculationOutput, error) {
	operator, err := valueobject.NewOperator(input.Operator)
	if err != nil {
		return dto.CalculationOutput{}, err
	}

	leftOperand := valueobject.NewOperand(input.LeftOperand)
	rightOperand := valueobject.NewOperand(input.RightOperand)

	operation := entity.NewOperation(leftOperand, operator, rightOperand)
	calculation := aggregate.NewCalculation(*operation)

	result, err := calculation.Compute()
	if err != nil {
		return dto.CalculationOutput{}, err
	}

	return dto.CalculationOutput{Result: result}, nil
}
