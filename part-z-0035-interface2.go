package main

import (
	"fmt"
)

func main() {

	rect := Rectangle {width: 12, height: 17}

	circ := Circle {radius: 19}

	fmt.Println(rect.area(), rect.perimeter())

	fmt.Println(circ.area(), circ.perimeter())
}

type Geo interface {

	area() uint

	perimeter() uint
}

type Rectangle struct {

	width, height uint
}

func (r Rectangle) area() uint {

	return r.height * r.width
}

func (r Rectangle) perimeter() uint {

	return 2 * (r.height + r.width)
}

type Circle struct {

	radius uint
}

func (c Circle) area() uint {

	return 3 * c.radius * c. radius
}

func (c Circle) perimeter() uint {

	return 6 * c.radius
}

// 204 58
// 1083 114
