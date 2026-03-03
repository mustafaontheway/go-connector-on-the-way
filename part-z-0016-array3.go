package main

import (
	"fmt"
)

func main() {

	var years = [4]uint16 {2000, 2001, 2002, 2003}

	for i := uint8(0); i < uint8(len(years)); i++ {

		fmt.Printf("Year%d: %d\n", i + 1, years[i])
	}
}

// Year1: 2000
// Year2: 2001
// Year3: 2002
// Year4: 2003
