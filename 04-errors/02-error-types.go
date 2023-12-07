package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	var multiplier, divisor int
	for {
		fmt.Println("Enter the multiplier and divisor")
		fmt.Scanln(&multiplier, &divisor)
		q, r, err := divide(multiplier, divisor)
		if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by 0")
			continue
		}
		if err != nil {
			fmt.Println("something went wrong :", err)
			break
		}
		fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
		break
	}
}

// using named results
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
