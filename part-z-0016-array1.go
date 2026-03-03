package main

import (
	"fmt"
)

func main() {

	var years = [4]uint16 {2000, 2001, 2002, 2003}

	var ages = [5]uint8 {17, 47, 57, 77}

	var rates = [...]float32 {3.21, -12.34, 43.21}

	fmt.Println(years)
	fmt.Println(ages)
	fmt.Println(rates)
}

// [2000 2001 2002 2003]
// [17 47 57 77 0]
// [3.21 -12.34 43.21]
