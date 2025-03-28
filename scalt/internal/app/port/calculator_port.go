package port

import "scalt/internal/app/dto"

type CalculatorPort interface {
	PerformCalculation(input dto.CalculationInput) (dto.CalculationOutput, error)
}
