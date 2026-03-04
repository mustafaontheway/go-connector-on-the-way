package main

import (
	"fmt"
)

func main() {

	yearsArray := [3]uint16 {1990, 2000, 2010}

	years := yearsArray[:]

	years = append(years, 2020, 2030)

	fmt.Println(years) // [1990 2000 2010 2020 2030]

	birthYears := []uint16 {1850, 1900, 1950}

	years = append(years, birthYears...)

	fmt.Println(years) // [1990 2000 2010 2020 2030 1850 1900 1950]
}
