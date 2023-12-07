package calculator

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func init() {
	fmt.Println("[calculator - divide.go] init invoked")
}

func Divide(x, y int) (quotient, remainder int, err error) {
	opCount["Divide"]++
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
