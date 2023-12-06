/*
Refactor the below program using functions
Make sure that every function has ONLY ONE reason to change (SRP)
*/

package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		fmt.Printf("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit\n")
		fmt.Println("Enter your choice")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
