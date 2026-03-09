package main

import (
	"fmt"
)

func main() {

	var base, pow, result uint 

	base = 4
	pow = 5
	result = 1
	
	for i := uint(0); i < pow; i++ {

		result *= base
	}

	fmt.Println(result) // 1024
}

