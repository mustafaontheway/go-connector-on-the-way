package main

import (
	"fmt"
)

func main() {

}

func findMax(x, y, z int) int {

	if x >= y && x >= z {

		return x
	} 

	if y >= x && y >= z {

		return y
	}

	return z
}

func findMin(x, y, z int) int {

	if x <= y && x <= z {

		return x
	} 

	if y <= x && y <= z {

		return y
	}
	
	return z
}

