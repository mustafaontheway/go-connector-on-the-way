package main

import (
	"fmt"
)

func main() {

	calculus(12, 76)
}

func calculus(x, y int) {

	var sum = func() {

		fmt.Printf("%d + %d = %d\n", x, y, x + y)
	}

	var diff = func() {

		fmt.Printf("%d - %d = %d\n", x, y, x - y)
	}

	var mult = func() {

		fmt.Printf("%d * %d = %d\n", x, y, x * y)
	}

	sum()
	diff()
	mult()
}

/*
12 + 76 = 88
12 - 76 = -64
12 * 76 = 912

*/
