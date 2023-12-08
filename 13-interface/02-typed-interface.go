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
func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area :", x.Area())
}
*/

type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

func (t Triangle) Perimeter() float32 {
	return t.Base + float32(math.Sqrt((float64((t.Base * t.Base) + (4 * t.Height * t.Height)))))
}

/*
	func PrintPerimeter(x interface{ Perimeter() float32 }) {
		fmt.Println("Perimeter :", x.Perimeter())
	}
*/
type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// v5.0 (interface composition)
/*
func PrintStats(x interface {
	interface{ Area() float32 }
	interface{ Perimeter() float32 }
}) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}
*/

/*
func PrintStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}
*/

type StatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x StatsFinder) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

	t := Triangle{Base: 20, Height: 10}
	// fmt.Println("Area :", t.Area())
	/*
		PrintArea(t)
		PrintPerimeter(t)
	*/
	PrintStats(t)

}
