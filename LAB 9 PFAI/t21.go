package main

import (
	"fmt"
	"math"
)

// Shape interface with Area method
type Shape interface {
	Area() float64
}

// Circle implements Shape
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle implements Shape
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// Create a slice of different shapes
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
		Circle{radius: 3},
		Rectangle{width: 10, height: 2},
	}

	// Calculate and print each shape's area
	for i, shape := range shapes {
		fmt.Printf("Shape %d area: %.2f\n", i+1, shape.Area())
	}
}
