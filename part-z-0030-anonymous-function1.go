package main

import (
	"fmt"
)

func main() {

	var warning = func() {

		fmt.Println("Be careful!")
	}

	warning()
	warning()
}

// Be careful!
// Be careful!
