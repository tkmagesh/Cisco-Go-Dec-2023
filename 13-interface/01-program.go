package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

// v2.0
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch o := x.(type) {
	case Circle:
		fmt.Println("Area :", o.Area())
	case Rectangle:
		fmt.Println("Area :", o.Area())
	default:
		fmt.Println("Invalid type")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch o := x.(type) {
	case interface{ Area() float32 }: //any object with Area() method
		fmt.Println("Area :", o.Area())
	default:
		fmt.Println("Invalid type")
	}
}
*/

func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area :", x.Area())
}

// v3.0
type Triangle struct {
	Base   float32
	Height float32
}

func (t Triangle) Area() float32 {
	return (t.Base * t.Height) / 2
}

// v4.0
// implement Perimeter() float32 for all the types and Print the perimeter
/*
	Perimeter formula:
		Cirle = 2 * Pi * r
		Rectangle = 2 * (H + W)
		Triangle (𝑏+ sqrt(𝑏*b + 4(ℎ*h)))
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	t := Triangle{Base: 20, Height: 10}
	// fmt.Println("Area :", t.Area())
	PrintArea(t)

	// PrintArea(100)
}
