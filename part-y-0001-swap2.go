package main

import (
	"fmt"
)

func main() {

	var a, b uint8 

	fmt.Println("Type your age:")

	fmt.Scan(&a)

	fmt.Println("Type his age:")

	fmt.Scan(&b)

	a, b = b, a
	
	fmt.Println("Your new age:", a)
	fmt.Println("His new age:", b)
}

// Type your age:
// 99
// Type his age:
// 19
// Your new age: 19
// His new age: 99
