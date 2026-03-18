package main

import (
	"fmt"
)

func main() {

	myNum := 17

	plus1(myNum)

	fmt.Println("My num is:", myNum) // My num is: 17

	fmt.Println("Address of myNum:", &myNum) // Address of myNum: 0xc8e060

	plus2(&myNum)

	fmt.Println("My num is:", myNum) // My num is: 18

	fmt.Println("Address of myNum:", &myNum) // Address of myNum: 0x108e060
}

func plus1(num int) {

	num += 1
}

func plus2(num *int) {

	*num += 1
}




