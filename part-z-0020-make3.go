package main

import (
	"fmt"
)

func main() {

	nums := [...]uint8 {12, 47, 87, 50, 99, 64, 47, 88}

	lenN := len(nums)

	evens := make([]uint8, 0, lenN)

	odds := make([]uint8, 0, lenN)

	for i := 0; i < lenN; i++ {

		if nums[i] % 2 == 0 {

			evens = append(evens, nums[i])

		} else {

			odds = append(odds, nums[i])
		}
	}

	fmt.Println("Even numbers:", evens)

	fmt.Println("Odd numbers:", odds)
}

// Even numbers: [12 50 64 88]
// Odd numbers: [47 87 99 47]
