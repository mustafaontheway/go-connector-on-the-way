package main
import "fmt";

func main() {

	r := sum(6, 12, -500)

	fmt.Println(r + 5000) // 4518
}

func sum(x int, y int, z int) int {

	return x + y + z
}
