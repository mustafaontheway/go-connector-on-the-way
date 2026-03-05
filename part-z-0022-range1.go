package main

import (
	"fmt"
)

func main() {

	years := [5]uint16 {1900, 1920, 1940, 1960, 1980}

	for i, v := range years {

		fmt.Printf("Year%d: %d\n", i + 1, v)
	}
}

// Year1: 1900
// Year2: 1920
// Year3: 1940
// Year4: 1960
// Year5: 1980
