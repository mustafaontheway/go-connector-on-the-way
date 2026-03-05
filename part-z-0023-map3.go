package main

import (
	"fmt"
)

func main() {

	employees := map[uint8] string {47: "ayhan bilir", 59: "bengü güzel"}

	employees[68] = "hakan burada"

	employees[47] = "aykan bilir"

	delete(employees, 59)

	fmt.Println(employees)
}

// map[47:aykan bilir 68:hakan burada]
