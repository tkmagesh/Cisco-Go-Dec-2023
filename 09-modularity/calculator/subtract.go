package calculator

import "fmt"

func init() {
	fmt.Println("[calculator - subtract.go] init invoked")
}

func Subtract(x, y int) int {
	opCount["Subtract"]++
	return x - y
}
