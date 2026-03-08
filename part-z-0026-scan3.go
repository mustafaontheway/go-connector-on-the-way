package main

import (
	"fmt"
)

func main() {

	var year uint16 
	var age uint8

	fmt.Println("Type current year and age value (use numeric and positive values):")

	fmt.Scan(&year, &age)

	fmt.Printf("Your birth year is: %d.\n", year - uint16(age))
}

