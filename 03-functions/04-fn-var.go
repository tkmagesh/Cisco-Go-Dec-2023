package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	addResult := add(100, 200)
	fmt.Println(addResult)

	// named results
	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
