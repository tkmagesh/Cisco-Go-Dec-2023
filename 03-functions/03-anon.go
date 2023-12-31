package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(addResult)

	// named results
	q, r := func(x, y int) (quotient, remainder int) {
		/*
			quotient = x / y
			remainder = x % y
		*/
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
