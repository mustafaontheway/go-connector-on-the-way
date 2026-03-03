package main

import (
	"fmt"
)

func main() {

	var num int8 

	for i := int8(1); i <= 20; i += 2 {

		num += i 

		if i == 7 {

			break
		}
	}

	fmt.Println(num) // 16
}
