package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Ayhan Ah"

	newName := strings.Replace(name, "h", "k", 1)

	fmt.Println(newName) // Aykan Ah

	name2 := "Aykan Hanak"

	newName2 := strings.ReplaceAll(name2, "k", "h")

	fmt.Println(newName2) // Ayhan Hanah
}




// go run main.go
