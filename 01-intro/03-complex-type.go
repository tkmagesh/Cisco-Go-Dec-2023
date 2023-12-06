package main

import "fmt"

func main() {
	/*
		var c1 complex128
		c1 = 3 + 5i
	*/
	c1 := 3 + 5i
	fmt.Println(c1)

	c2 := 5 + 7i
	fmt.Println(c2)

	x := c1 + c2
	fmt.Println(x)

	fmt.Println("real(x) :", real(x))
	fmt.Println("imag(x) :", imag(x))
}
