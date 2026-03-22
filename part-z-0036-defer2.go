package main

import (
	"fmt"
)

func main() {

	// defer
	One()
}

func One() {

	fmt.Println("Hi 1")
	fmt.Println("Hi 2")
	defer fmt.Println("Hi 3")
	fmt.Println("Hi 4")
	fmt.Println("Hi 5")
}

// Hi 1
// Hi 2
// Hi 4
// Hi 5
// Hi 3
