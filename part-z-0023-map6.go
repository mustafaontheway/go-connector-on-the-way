package main

import (
	"fmt"
)

func main() {

	empSalaries := make(map[string] uint16)

	empSalaries["au004296"] = 4300
	empSalaries["kt009654"] = 3700
	empSalaries["mn004937"] = 4400

	searchSalary, hasSal := empSalaries["mn004937"]

	fmt.Println(searchSalary)
	fmt.Println(hasSal)
}

// 4400
// true
