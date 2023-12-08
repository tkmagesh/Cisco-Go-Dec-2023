package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product Product) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
}

func (product *Product) applyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	var pen Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 5,
	}
	// PrintProduct(pen)
	pen.Print()
	pen.applyDiscount(10)
	pen.Print()

	// accessing the methods using pointer
	var pencilPtr = &Product{101, "Pencil", 2}
	pencilPtr.Print()
	pencilPtr.applyDiscount(20)
	pencilPtr.Print()

}
