package main

import (
	"fmt"
)

func main() {

	var ages = []uint8 {18, 69, 87, 65, 88, 99}

	updateAges(ages)

	fmt.Println(ages) // [17 69 87 65 88 17]
}

func updateAges(slc []uint8) {

	slc[0] = 17
	slc[len(slc) - 1] = 17
}




