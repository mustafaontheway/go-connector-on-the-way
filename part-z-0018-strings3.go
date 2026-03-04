package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Mustafa Bilge Buyukdereli."

	splitted1 := strings.Split(name, " ")

	fmt.Println(splitted1) // [Mustafa Bilge Buyukdereli.]

	splitted2 := strings.Split(strings.TrimSuffix(name, "."), " ")

	fmt.Println(splitted2) // [Mustafa Bilge Buyukdereli]

	fmt.Println(splitted2[2]) // Buyukdereli
}
