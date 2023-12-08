package main

import "fmt"

func main() {
	// var x interface{} // before go 1.18
	var x any // from go 1.18
	x = 100
	x = true
	x = "Dolore irure do ipsum est excepteur sunt tempor non pariatur in Lorem ad ut nulla."
	x = 91.99
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	/*
		x = "Elit veniam aliquip esse consequat incididunt aute ipsum occaecat irure duis."
		y := x.(int) + 200 //unsafe type conversion (possibility of a runtime panic)
		fmt.Println(y)
	*/

	x = 100
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = "In ut commodo dolor fugiat elit nulla dolor in nisi aliquip veniam cillum."
	// x = true
	// x = 99.99
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 0.01% of x =", val*(0.01/100))
	default:
		fmt.Println("Unknown type")

	}

}
