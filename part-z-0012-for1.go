package main

import (
	"fmt"
)

func main() {

	var year uint16 = 1900

	for i := 1; i < 6; i++ {

		fmt.Printf("Year_%d : %d\n", i, year)

		year++
	} 
}

/*
Year_1 : 1900
Year_2 : 1901
Year_3 : 1902
Year_4 : 1903
Year_5 : 1904

*/
