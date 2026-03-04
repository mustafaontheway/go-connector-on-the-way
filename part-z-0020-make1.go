package main

import (
	"fmt"
)

func main() {

	yearsSlice1 := make([]uint16, 0, 5)

	fmt.Println(len(yearsSlice1), cap(yearsSlice1)) // 0 5

	fmt.Println(yearsSlice1) // []

	yearsSlice2 := make([]uint16, 3, 5)

	fmt.Println(len(yearsSlice2), cap(yearsSlice2)) // 3 5

	fmt.Println(yearsSlice2) // [0 0 0]
}
