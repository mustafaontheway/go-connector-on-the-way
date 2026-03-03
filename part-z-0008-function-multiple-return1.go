package main
import "fmt";

func main() {

	s, d := sumOrDiff(6, 12)

	fmt.Println("Sum:", s)

	fmt.Println("Diff:", d)
}

func sumOrDiff(x int, y int) (int, int) {

	return x + y, x - y
}

// Sum: 18
// Diff: -6
