package main

import (
	"fmt"
)

func main() {

	fmt.Println(sumNums(-10, 7)) // -27
}

func sumNums(firstNum, lastNum int) int {

	sum := 0

	for firstNum <= lastNum {

		sum += firstNum

		firstNum++
	}

	return sum
}
