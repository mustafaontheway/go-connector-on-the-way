package main

import (
	"fmt"
)

func main() {

	yearsSlice1 := make([]uint16, 0, 5)

	yearsSlice1 = append(yearsSlice1, 1000, 2000, 30, 40, 50, 60, 70)

	fmt.Println(yearsSlice1) // [1000 2000 30 40 50 60 70]

	fmt.Println(len(yearsSlice1), cap(yearsSlice1)) // 7 12

	yearsSlice2 := make([]uint16, 3, 5)

	yearsSlice2 = append(yearsSlice2, 2000, 2010, 2020)

	fmt.Println(yearsSlice2) // [0 0 0 2000 2010 2020]

	fmt.Println(len(yearsSlice2), cap(yearsSlice2)) // 6 12
}
