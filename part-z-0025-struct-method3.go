package main

import (
	"fmt"
)

func main() {

	empAyhan := Employee{"Ayhan Bilir", "HR", "ab009874", 4400}

	empAyhan.updateSalary(3000)

	fmt.Println(empAyhan.salaryUSD)
}

type Employee struct {

	name string
	department string
	empID string
	salaryUSD uint16
}

func (e *Employee) updateSalary(newSalaryUSD uint16) {

	e.salaryUSD = newSalaryUSD
}




