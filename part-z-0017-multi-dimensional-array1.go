package main

import (
	"fmt"
)

func main() {

	var results = [3][4]int8 {{-12, 21, 10, 4}, {17, -5, 9, 14}, {21, -100, 45, 14}}

	fmt.Println("Min number of results array:", results[2][1]) // Min number of results array: -100
}
