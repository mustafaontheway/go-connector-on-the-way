package main

import (
	"fmt"
)

func main() {

	ages := []uint8 {37, 21, 96, 87, 15, 56, 74, 5}

	maxAge := ages[0]

	minAge := ages[0]

	for _, age := range ages {

		if age > maxAge {

			maxAge = age
		}

		if age < minAge {

			minAge = age
		}
	}

	fmt.Println("Max age:", maxAge) // Max age: 96

	fmt.Println("Min age:", minAge) // Min age: 5
}

