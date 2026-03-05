package main

import (
	"fmt"
)

type Shape struct {

	width float64
	height float64
}

func (s Shape) calculateArea() float64 {

	return s.height * s.width
}

func (s Shape) calculatePerimeter() {

	fmt.Printf("Width: %.2f - height: %.2f -> perimeter: %.2f", s.width, s.height, 2 * (s.height + s.width))
}

func main() {

	shape1 := Shape {width: 3.24, height: 4}

	area := shape1.calculateArea()

	fmt.Println(area)

	shape1.calculatePerimeter()
}

// 12.96
// Width: 3.24 - height: 4.00 -> perimeter: 14.48
