package calculator

import "fmt"

var opCount map[string]int

func init() {
	opCount = make(map[string]int)
	fmt.Println("[calculator - calc.go] init invoked")
}

func OpCount() map[string]int {
	return opCount
}

func Multiply(x, y int) int {
	opCount["Multiply"]++
	return x * y
}
