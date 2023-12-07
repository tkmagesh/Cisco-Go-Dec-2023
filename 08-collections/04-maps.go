package main

import "fmt"

func main() {
	// var productRanks map[string]int

	// creation & initialization
	// var productRanks map[string]int = map[string]int{"pen": 3, "marker": 5, "pencil": 2}
	// var productRanks = map[string]int{"pen": 3, "marker": 5, "pencil": 2}
	// productRanks := map[string]int{"pen": 3, "marker": 5, "pencil": 2}
	/*
		productRanks := map[string]int{
			"pen":    3,
			"marker": 5,
			"pencil": 2, // note the comma even on the last item
		}
	*/

	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
	*/

	// var productRanks map[string]int = make(map[string]int)
	var productRanks map[string]int = map[string]int{}
	productRanks["pen"] = 3
	productRanks["marker"] = 5
	productRanks["pencil"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating a map using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Removing an item")
	keyToRemove := "notebook"
	delete(productRanks, keyToRemove) // noop if the key doesn't exist
	fmt.Println(productRanks)

	fmt.Println("Check if an item exists")
	keyToCheck := "notebook"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Product %q doesn't exist\n", keyToCheck)
	}

}
