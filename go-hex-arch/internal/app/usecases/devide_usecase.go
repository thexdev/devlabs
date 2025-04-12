package usecases

import (
	"gohexarch/internal/domain/models"
)

type DevideUseCase struct {
}

func (uc *DevideUseCase) Execute(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Devide(a, b)
}
