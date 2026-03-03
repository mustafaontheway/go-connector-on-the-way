package main
import "fmt";

func main() {

	x := 10
	y := 7

	var cond string

	if x > y && y * y >= 50 {

		cond = "Condition 1"

	} else if x + y > 50 ||  x % y > 1 {

		cond = "Condition 2"

	} else {

		cond = "Condition 3"
	}

	fmt.Println(cond) // Condition 2
}

// go run main.go
