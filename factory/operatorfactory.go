package factory

import (
	"calculator-svc/operation"
	"errors"
)

var ErrInvalidOperator = errors.New("invalid operator")

type OperatorFactory struct{}

func (o OperatorFactory) Factory(s string) (operation.Operator, error) {
	switch s {
	case "+":
		return operation.NewAddition(), nil
	case "-":
		return operation.NewSubtraction(), nil
	case "*":
		return operation.NewMultiplication(), nil
	case "/":
		return operation.NewDivision(), nil
	default:
		return nil, ErrInvalidOperator
	}
}
