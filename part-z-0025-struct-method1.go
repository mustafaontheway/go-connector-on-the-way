package main

import (
	"fmt"
)

type Student struct {

	name string
	university string
	age uint8
}

func (s Student) printInfo() {

	fmt.Printf("Student name: %s - university: %s - age: %d\n", s.name, s.university, s.age)
}

func main() {

	studentAyhan := Student {"Ayhan Bilir", "İstiklal University", 23}

	studentAyhan.printInfo()

	// Student name: Ayhan Bilir - university: İstiklal University - age: 23
}

