package main

import (
	"fmt"
)

func main() {

	type Student struct {

		name string
		university string
		age uint8
	}

	studentAyhan := Student {"Ayhan Bilir", "İstiklal University", 23}

	fmt.Println(studentAyhan)
	fmt.Println(studentAyhan.name)
	fmt.Println(studentAyhan.university)
	fmt.Println(studentAyhan.age)
}

// {Ayhan Bilir İstiklal University 23}
// Ayhan Bilir        
// İstiklal University
// 23
