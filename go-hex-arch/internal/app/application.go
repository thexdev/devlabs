package app

import (
	"gohexarch/internal/domain/models"
)

type Application struct {
	A float32
	B float32
}

func (app *Application) Add(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Add(a, b)
}

func (app *Application) Subtract(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Subtract(a, b)
}

func (app *Application) Devide(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Devide(a, b)
}

func (app *Application) Multiply(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)
	return calc.Multiply(a, b)
}
