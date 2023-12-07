package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("App panicked. err :", err)
		}
		fmt.Println("Thank you!")
	}()
	for {
		var multiplier, divisor int
		fmt.Println("Enter the multiplier and divisor")
		inputCount, err := fmt.Scanf("%d %d\n", &multiplier, &divisor)
		if inputCount != 2 {
			fmt.Println("Invalid input :", err)
			continue
		}
		q, r, err := divideClient(multiplier, divisor)
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

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	remainder = x % y
	return
}
