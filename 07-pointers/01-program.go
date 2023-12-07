package main

import "fmt"

func main() {
	var no int = 100

	var noPtr *int
	noPtr = &no

	// accessing the data using the pointer (deferencing)
	fmt.Println(noPtr, *noPtr)

	fmt.Println("Before incrementing, no :", no)
	increment(&no)
	fmt.Println("After incrementing, no :", no)

	x, y := 100, 200
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap( /* ?? */ )
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(x *int) {
	fmt.Println("x :", x)
	*x++
}

func swap( /* ? */ ) /* no return values */ {

}
