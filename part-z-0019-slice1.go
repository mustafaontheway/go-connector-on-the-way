package main

import (
	"fmt"
)

func main() {

	nums := []int8 {3, -33, 22} // slice

	nums = append(nums, -13)

	fmt.Println(nums) // [3 -33 22 -13]

	yearsArray := [3]uint16 {1990, 2000, 2010}

	years := yearsArray[:]

	years = append(years, 2020, 2030)

	fmt.Println(years) // [1990 2000 2010 2020 2030]
}

