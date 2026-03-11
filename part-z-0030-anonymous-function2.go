package main

import (
	"fmt"
)

func main() {

	var warning = func(name string) {

		fmt.Printf("%s, be careful!\n", name)
	}

	warning("Ayhan")
	warning("Aykan")
}

// Ayhan, be careful!
// Aykan, be careful!
