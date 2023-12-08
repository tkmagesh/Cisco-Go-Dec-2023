package main

import "fmt"

func main() {
	// var nos []int = []int{10, 20, 30}
	/*
		var product struct {
			Id   int
			Name string
			Cost float32
		}
	*/

	var pen struct {
		Id   int
		Name string
		Cost float32
	} = struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   100,
		Name: "Pen",
		Cost: 5,
	}
	// fmt.Printf("%#v\n", pen)
	PrintProduct(pen)
}

func PrintProduct(product struct {
	Id   int
	Name string
	Cost float32
}) {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
}
