package main

import (
	"fmt"
)

func main() {

	empSalaries := make(map[string] uint16)

	empSalaries["au004296"] = 4300
	empSalaries["kt009654"] = 3700
	empSalaries["mn004937"] = 4400

	fmt.Println(empSalaries) // map[au004296:4300 kt009654:3700 mn004937:4400]
}
