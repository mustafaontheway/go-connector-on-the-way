package main

// import (
// 	"fmt"
// )

func main() {


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
