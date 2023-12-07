package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	// nos := [5]int{3, 1, 4}
	fmt.Println(nos)

	fmt.Println("Iteration using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iteration using range")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}

	/*
		nos2 := nos // creating a copy of the array
		nos2[0] = 9999
		fmt.Printf("nos : %v, nos2 : %v\n", nos, nos2)
	*/

	nosPtr := &nos
	// fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0]) // members can be accessed from the pointer as well

	// use the sort function to sort the nos array
	fmt.Println("After sorting:")
	sort(&nos)
	fmt.Println(nos)
}

func sort(values *[5]int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
