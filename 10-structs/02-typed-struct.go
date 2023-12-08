package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	// var nos []int = []int{10, 20, 30}
	/*
		var product Product
	*/

	// var pen Product = Product{100, "Pen", 5}
	// var pen Product = Product{100, "Pen"} // cannot partially initialize
	/*
		var pen Product = Product{
			Id:   100,
			Name: "Pen",
		}
	*/

	/*
		var pen Product
		pen.Id = 100
		pen.Name = "Pen"
		pen.Cost = 5
	*/

	// the below is a standard approach
	var pen Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 5,
	}
	// fmt.Printf("%#v\n", pen)
	PrintProduct(pen)
	fountainPen := pen // creates a copy
	fountainPen.Cost = 10

	// pointers
	var penPtr *Product
	penPtr = &pen
	fmt.Println(penPtr.Id) // members can be accessed using the "." notation from pointers as well

	fmt.Println("After applying 10% discount")
	applyDiscount(&pen, 10)
	PrintProduct(pen)

}

func PrintProduct(product Product) {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
}

func applyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
