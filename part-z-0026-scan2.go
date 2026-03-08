package main

import (
	"fmt"
)

func main() {

	var year uint16 
	var age uint8

	fmt.Println("Type current year:")

	fmt.Scan(&year)

	fmt.Println("Type your age:")

	fmt.Scan(&age)

	fmt.Printf("Your birth year is: %d.\n", year - uint16(age))
}

