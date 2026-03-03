package main
import "fmt";

func main() {

	var x int8 = -5 

	var y int8 = 10 

	var cond1 bool = x != y && -x < y 

	var cond2 bool = x >= y || x * x == y 

	fmt.Println(cond1, cond2) // true false
}
