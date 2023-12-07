package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {

	for {
		var multiplier, divisor int
		fmt.Println("Enter the multiplier and divisor")
		inputCount, err := fmt.Scanf("%d %d\n", &multiplier, &divisor)
		if inputCount != 2 {
			fmt.Println("Invalid input :", err)
			continue
		}
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
	fmt.Printf("[divide] x = %d and y = %d\n", x, y)
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
