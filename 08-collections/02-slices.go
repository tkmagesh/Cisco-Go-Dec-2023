package main

import "fmt"

func main() {

	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// adding new items
	nos = append(nos, 10)
	nos = append(nos, 20, 30)
	fmt.Println(nos)

	// append another slice
	hundreds := []int{100, 200}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Println("Iteration using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iteration using range")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}

	// slicing
	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:5] =", nos[:5])

	/*
		nos2 := nos
		nos2[0] = 9999
		fmt.Printf("nos = %v, nos2 = %v\n", nos, nos2)
	*/

	s2 := nos[2:5]
	s2[0] = 9999
	fmt.Printf("nos = %v, s2 = %v\n", nos, s2)
	fmt.Printf("cap(nos) = %v, cap(s2) = %v\n", cap(nos), cap(s2))

	fmt.Println("After sorting")
	sort(nos)
	fmt.Println(nos)
}

func sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
