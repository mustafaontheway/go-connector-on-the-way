package main

import (
	"fmt"
)

func main() {

	nums := []uint {12, 47, 87, 50, 99, 64, 47, 88}

	n1 := make([]uint, len(nums))

	copy(n1, nums)

	fmt.Println("New copied slice (version 1):", n1)

	n1[0] = 50_000

	fmt.Println("New copied slice (version 2):", n1)

	fmt.Println("Original slice:", nums)
}

// New copied slice (version 1): [12 47 87 50 99 64 47 88]
// New copied slice (version 2): [50000 47 87 50 99 64 47 88]
// Original slice: [12 47 87 50 99 64 47 88]
