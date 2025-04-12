package usecases

import "gohexarch/internal/domain/models"

type SubstractUseCase struct {
}

func (uc *SubstractUseCase) Execute(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Subtract(a, b)
}
