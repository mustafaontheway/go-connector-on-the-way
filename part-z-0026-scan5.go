package main

import (
	"fmt"
)

func main() {

	var a, b, c, d uint8 
	var finalNum uint16

	fmt.Println("Please, create number abcd (5762, 8964, 9438 etc.). Type number d:")

	fmt.Scan(&d)

	fmt.Println("Type number c:")

	fmt.Scan(&c)

	fmt.Println("Type number b:")

	fmt.Scan(&b)

	fmt.Println("Type number a:")

	fmt.Scan(&a)

	finalNum = 1000 * uint16(a) + 100 * uint16(b) + 10 * uint16(c) + uint16(d)
	
	fmt.Println("Our number is now:", finalNum)
}

/*

Please, create number abcd (5762, 8964, 9438 etc.). Type number d:
3
Type number c:
1
Type number b:
9
Type number a:
8
Our number is now: 8913

*/
