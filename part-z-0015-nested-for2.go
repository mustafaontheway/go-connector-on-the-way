package main

import (
	"fmt"
)

func main() {

	var results = [3][4]int8 {{-12, 21, 10, 4}, {17, -5, 9, 14}, {21, -100, 45, 14}}

	fmt.Println(results) // [[-12 21 10 4] [17 -5 9 14] [21 -100 45 14]]

	var negTotal int16
	var posTotal uint16 

	for i := uint8(0); i < 3; i++ {

		for j := uint8(0); j < 4; j++ {

			if results[i][j] > 0 {

				posTotal += uint16(results[i][j])

			} else {

				negTotal += int16(results[i][j])
			}
		}
	}

	fmt.Println("Positive total:", posTotal)

	fmt.Println("Negative total:", negTotal)
}

// Positive total: 155
// Negative total: -117


// go run main.go
