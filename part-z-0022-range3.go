package main

import (
	"fmt"
)

func main() {

	ages := []uint8 {37, 21, 96, 87, 15, 96, 74, 5}

	var totalAges uint

	for _, age := range ages {

		totalAges += uint(age)
	}

	fmt.Println("Average age:", totalAges/ uint(len(ages))) // Average age: 53

	fmt.Println("Average age:", float32(totalAges)/ float32(len(ages))) // Average age: 53.875
}

