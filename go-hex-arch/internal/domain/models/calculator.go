package models

import "errors"

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a float32, b float32) (float32, error) {
	return a + b, nil
}

func (c *Calculator) Subtract(a float32, b float32) (float32, error) {
	return a - b, nil
}

func (c *Calculator) Multiply(a float32, b float32) (float32, error) {
	return a * b, nil
}

func (c *Calculator) Devide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
