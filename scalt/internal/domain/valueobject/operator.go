package valueobject

import "errors"

type Operator string

const (
	Addition       Operator = "+"
	Subtraction    Operator = "-"
	Multiplication Operator = "*"
	Division       Operator = "/"
)

func NewOperator(symbol string) (Operator, error) {
	switch symbol {
	case "+", "-", "*", "/":
		return Operator(symbol), nil
	default:
		return "", errors.New("invalid operator")
	}
}

func (o Operator) Symbol() string {
	return string(o)
}
