package main

import (
	"fmt"
)

func main() {

	fmt.Println(sumNumbers(14, 75, 999))

	fmt.Println(sumNumbers(23.3414, -75.22222222222222, 444.3))

}

func sumNumbers(nums ...float64) float64 {

	var sum float64 = 0

	for i := 0; i < len(nums); i++ {

		sum += nums[i]
	}

	return sum
}

// 1088
// 392.4191777777778
