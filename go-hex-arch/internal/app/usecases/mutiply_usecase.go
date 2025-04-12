package usecases

import "gohexarch/internal/domain/models"

type MultiplyUseCase struct {
}

func (uc *MultiplyUseCase) Execute(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Multiply(a, b)
}
