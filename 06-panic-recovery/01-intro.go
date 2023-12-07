package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("App panicked. err [1] :", err1)
		}
		if err2 := recover(); err2 != nil {
			fmt.Println("App panicked. err [2] :", err2)
		}
		fmt.Println("Thank you!")
	}()
	var multiplier, divisor int
	fmt.Println("Enter the multiplier and divisor")
	fmt.Scanln(&multiplier, &divisor)
	q, r := divide(multiplier, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
}

// using named results
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	remainder = x % y
	return
}
