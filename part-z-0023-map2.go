package main

import (
	"fmt"
)

func main() {

	employees := map[uint8] string {47: "ayhan bilir", 59: "bengü güzel"}

	employees[68] = "hakan burada"

	fmt.Println(employees)

	fmt.Println(employees[59])
}

// map[47:ayhan bilir 59:bengü güzel 68:hakan burada]
// bengü güzel
