package main

import (
	"fmt"
)

func main() {

	var years = [4]uint16 {2000, 2001, 2002, 2003}

	fmt.Println(years[2])

	var ages [5]uint8 

	ages[0] = 17
	ages[1] = 47
	ages[2] = 37
	ages[3] = 67
	ages[4] = 97

	fmt.Println(len(ages))
}

// 2002
// 5


// go run main.go
