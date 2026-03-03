package main

import (
	"fmt"
)

func main() {

	var num int8 

	for i := int8(1); i <= 20; i += 2 {

		if i == 9 {

			continue
		}
			
		num += i 
	}

	fmt.Println(num) // 91
}
