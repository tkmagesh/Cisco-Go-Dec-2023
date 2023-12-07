package main

import (
	"errors"
	"fmt"
)

func main() {
	var multiplier, divisor int
	for {
		fmt.Println("Enter the multiplier and divisor")
		fmt.Scanln(&multiplier, &divisor)
		q, r, err := divide(multiplier, divisor)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
		break
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

// using named results
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
