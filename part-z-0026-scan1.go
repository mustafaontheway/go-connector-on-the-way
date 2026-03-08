package main

import (
	"fmt"
)

func main() {

	var year uint16 

	fmt.Println("Type current year (only positive...):")

	fmt.Scan(&year)

	fmt.Println("Last year is", year - 1)
}

