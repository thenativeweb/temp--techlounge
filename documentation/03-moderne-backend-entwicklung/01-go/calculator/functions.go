package calculator

import "errors"

var ErrDivisionByZero = errors.New("division by zero")

func Add(left float64, right float64) float64 {
	return left + right
}

func Subtract(left float64, right float64) float64 {
	return left - right
}

func Multiply(left float64, right float64) float64 {
	return left * right
}

func Divide(left float64, right float64) (float64, error) {
	if right == 0 {
		return 0, ErrDivisionByZero
	}

	return left / right, nil
}
