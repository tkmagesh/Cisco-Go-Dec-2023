package calculator

import "fmt"

func init() {
	fmt.Println("[calculator - add.go] init invoked")
}

func Add(x, y int) int {
	opCount["Add"]++
	return x + y
}
