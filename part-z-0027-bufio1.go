package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type a proverb:")

	p, _ := reader.ReadString('\n')

	fmt.Printf("%s", p)
}

// Type a proverb:
// A barking dog never bites
// A barking dog never bites
