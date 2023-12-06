package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
	fmt.Println(sum(0))
}

func sum(no int, nos ...int) int {
	result := no
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
