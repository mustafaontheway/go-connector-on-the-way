package main

import (
	"fmt"
)

func main() {

	// defer
	One()
}

// Func 1 came
// Func 3 came
// Func 2 came

func One() {

	fmt.Println("Func 1 came")
	defer Two()
	Three()
}

func Two() {

	fmt.Println("Func 2 came")
}

func Three() {

	fmt.Println("Func 3 came")
}
