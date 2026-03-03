package main

import (
	"fmt"
)

func main() {

	var month uint8 = 10

	switch month {

	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Autumn")
	default:
		fmt.Println("Month index should be between 1 and 12!")
	}
}
