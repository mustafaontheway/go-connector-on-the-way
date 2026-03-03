package main

import (
	"fmt"
)

func main() {

	var sumEvens int8 
	var sumOdds int8

	for index := int8(1); index <= 15; index++ {

		if index % 2 == 0 {

			sumEvens += index

		} else {

			sumOdds += index
		}
	} 

	fmt.Println("Sum of even numbers:", sumEvens)

	fmt.Println("Sum of odd numbers:", sumOdds)
}

// Sum of even numbers: 56
// Sum of odd numbers: 64
