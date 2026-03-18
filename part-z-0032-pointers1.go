package main

import (
	"fmt"
)

func main() {

	myNum := 17

	plus1(myNum)

	fmt.Println("My num is:", myNum) // My num is: 17

	plus2(&myNum)

	fmt.Println("My num is:", myNum) // My num is: 18
}

func plus1(num int) {

	num += 1
}

func plus2(num *int) {

	*num += 1
}




