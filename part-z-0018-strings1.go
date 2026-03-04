package main

import (
	"fmt"
	"strings"
)

func main() {

	name1 := "Ayhan"

	name2 := "ayhaN"

	var condition1= strings.Compare(strings.ToLower(name1), strings.ToLower(name2))

	fmt.Println(condition1) // 0

	//var condition2 = strings.ToLower(name1) == strings.ToLower(name2)

	var condition2 = strings.EqualFold(name1, name2)

	fmt.Println(condition2)  // true
}




// go run main.go
