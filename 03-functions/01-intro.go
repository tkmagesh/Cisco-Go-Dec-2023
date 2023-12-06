package main

import "fmt"

func main() {

	// no arguments & no return values
	sayHi()

	// 1 argument & no return values
	greet("Magesh")

	// 2 arguments & no return values
	greetUser("Magesh", "Kuppan")

	// 2 arguments & 1 return value
	fmt.Println(add(100, 200))

	// 2 arguments & 2 return values

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named results
func divide(x, y int) (quotient, remainder int) {
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
