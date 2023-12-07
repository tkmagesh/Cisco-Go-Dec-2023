/*
Refactor the below program using functions
Make sure that every function has ONLY ONE reason to change (SRP)
*/

package main

import (
	"errors"
	"fmt"
)

var ErrInvalidOperation error = errors.New("invalid operation")

func main() {
	var userChoice, n1, n2, result int
	for {
		userChoice = getUserChoice()

		if userChoice == 5 {
			break
		}

		oper, err := getOperation(userChoice)
		if err == ErrInvalidOperation {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 = getOperands()
		result = oper(n1, n2)
		fmt.Println("Result :", result)
	}
}

func getUserChoice() int {
	var userChoice int
	fmt.Printf("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit\n")
	fmt.Println("Enter your choice")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands")
	fmt.Scanln(&n1, &n2)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

// return an error if the userChoice is not 1 - 4
func getOperation(userChoice int) (op func(int, int) int, err error) {
	switch userChoice {
	case 1:
		op = add
	case 2:
		op = subtract
	case 3:
		op = multiply
	case 4:
		op = divide
	default:
		err = ErrInvalidOperation
	}
	return
}
