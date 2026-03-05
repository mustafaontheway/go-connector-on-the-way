package main

import (
	"fmt"
)

func main() {

	employees := map[uint8] string {47: "ayhan bilir", 59: "bengü güzel"}

	employees[68] = "hakan burada"

	for k, v := range employees {

		fmt.Printf("Employee ID: %d - name: %s\n", k, v)
	}
}

// Employee ID: 47 - name: ayhan bilir
// Employee ID: 59 - name: bengü güzel
// Employee ID: 68 - name: hakan burada
