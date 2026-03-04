package main

import (
	"fmt"
)

func main() {

	nums := []uint {12, 47, 87, 50, 99, 64, 47, 88}

	n1 := []uint {100, 200, 3000000}

	n2 := []uint {100, 200, 3000000, 500000000}

	copy(n1, nums)

	fmt.Println(nums)

	fmt.Println(n1)

	copy(n2, nums)

	fmt.Println(n2)
}

// [12 47 87 50 99 64 47 88]
// [12 47 87]
// [12 47 87 50]
