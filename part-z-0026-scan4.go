package main

import (
	"fmt"
)

func main() {

	var a, b, c int8

	fmt.Println("Type three integer numbers between -100 and 100:")

	fmt.Scan(&a, &b, &c)

	fmt.Printf("%d + %d + %d = %d\n", a, b, c, sumNums(a, b, c))
}

func sumNums(x, y, z int8) int16 {

	return int16(x) + int16(y) + int16(z)
}

// Type three integer numbers between -100 and 100:
// -32
// 45
// -99
// -32 + 45 + -99 = -86
