package main

import (
	"fmt"
)

func main() {

	empAyhan := newEmployee("Ayhan Bilir", "HR", "ab009874", 3400)

	fmt.Println(empAyhan) // {Ayhan Bilir HR ab009874 3400}
}

type Employee struct {

	name string
	department string
	empID string
	salaryUSD uint16
}

func newEmployee(name string, department string, empID string, salaryUSD uint16) *Employee {

	return &Employee{name, department, empID, salaryUSD}
}





