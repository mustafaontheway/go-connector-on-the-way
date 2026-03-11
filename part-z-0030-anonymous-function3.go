package main

import (
	"fmt"
)

func main() {

	var sumNums = func(x, y int) int {

		return x + y
	}

	fmt.Println(sumNums(100, -121)) // -21
}
