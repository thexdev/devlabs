package aggregate

import (
	"errors"
	"scalt/internal/domain/entity"
	"scalt/internal/domain/valueobject"
)

type Calculation struct {
	operation entity.Operation
	result    float64
}

func NewCalculation(operation entity.Operation) *Calculation {
	return &Calculation{
		operation: operation,
	}
}

func (c *Calculation) Compute() (float64, error) {
	lp := c.operation.LeftOperand.Value()
	rp := c.operation.RightOperand.Value()

	switch c.operation.Operator {
	case valueobject.Addition:
		c.result = lp + rp
	case valueobject.Subtraction:
		c.result = lp - rp
	case valueobject.Multiplication:
		c.result = lp * rp
	case valueobject.Division:
		if rp == 0 {
			return 0, errors.New("cannot devide by zero")
		}
		return lp / rp, nil
	default:
		return 0, errors.New("unsupported operator")
	}

	return c.result, nil
}

func (c *Calculation) Result() float64 {
	return c.result
}
