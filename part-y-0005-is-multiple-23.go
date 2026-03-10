package main

import (
	"fmt"
)

func main() {

	isMultipleOf_23(11, 196)
}

func isMultipleOf_23(first, last int) {

	for first <= last {

		if first % 23 == 0 {

			fmt.Printf("%d ", first)
		}
		first++
	}
}

// 23 46 69 92 115 138 161 184 
