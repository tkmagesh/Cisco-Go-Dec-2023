package main

import "fmt"

func main() {
	// var s []int
	// s := make([]int, 0, 10)
	s := make([]int, 7, 10)

	s = append(s, 10)
	fmt.Printf("len(s) = %d, cap(s) = %d, s = %v\n", len(s), cap(s), s)
	s = append(s, 20)
	fmt.Printf("len(s) = %d, cap(s) = %d, s = %v\n", len(s), cap(s), s)
	s = append(s, 30)
	fmt.Printf("len(s) = %d, cap(s) = %d, s = %v\n", len(s), cap(s), s)
	s = append(s, 40)
	fmt.Printf("len(s) = %d, cap(s) = %d, s = %v\n", len(s), cap(s), s)
	s = append(s, 50)
	fmt.Printf("len(s) = %d, cap(s) = %d, s = %v\n", len(s), cap(s), s)
}
