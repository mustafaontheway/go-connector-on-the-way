package main

import (
	"fmt"
)

func main() {

	a,b,c := findDigits(795)

	fmt.Println(a, b, c) // 5 9 7
}

func findDigits(num uint) (uint8, uint8, uint8) {

	if num < 100 || num > 999 {return 0, 0, 0}

	ones := num % 10 //795 % 10 -> 5
	tens := ((num - ones) / 10) % 10 //((795 - 5) / 10) % 10 -> 9
	hundreds := (num - (ones + tens * 10)) / 100 // (795 - (5 + 90)) / 100 -> 7

	return  uint8(ones), uint8(tens), uint8(hundreds)
}
