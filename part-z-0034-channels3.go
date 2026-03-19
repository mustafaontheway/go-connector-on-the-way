package main

import (
	"fmt"

)

func main() {

	evenCn := make(chan uint)

	oddCn := make(chan uint)

	fmt.Println("Started")

	fmt.Println("Hi 1")

	go sumEvens(250, evenCn)

	go sumOdds(600, oddCn)

	totalEven := <- evenCn

	totalOdd := <- oddCn

	fmt.Println(totalEven)

	fmt.Println("Hi 2")

	fmt.Println(totalOdd)

	fmt.Println("Ended")
}

func sumEvens(limit uint, evenCn chan uint) {

	var total uint

	for i := uint(0); i <= limit; i++ {

		if i % 2 == 0 {

			total += i
		}
	}

	evenCn <- total
}

func sumOdds(limit uint, oddCn chan uint) {

	var total uint 

	for i := uint(0); i <= limit; i++ {

		if i % 2 != 0 {

			total += i
		}
	}

	oddCn <- total
}

// Started
// Hi 1 
// 15750
// Hi 2 
// 90000
// Ended
