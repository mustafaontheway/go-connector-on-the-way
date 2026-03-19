package main

import (
	"fmt"

)

func main() {

	evenCn := make(chan uint)

	oddCn := make(chan uint)

	go sumEvens(18, evenCn)

	go sumOdds(20, oddCn)

	fmt.Println(evenCn)

	fmt.Println(oddCn)

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

// 0x1c162c0
// 0x1c16300
// Ended
