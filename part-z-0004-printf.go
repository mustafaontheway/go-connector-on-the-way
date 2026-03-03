package main
import "fmt";

func main() {

	/*
		integer	%d
		float	%g
		string	%s
		bool	%t
	*/

	var year uint16 = 2026 

	var expectedRate float32 = 99.99 

	name := "Mustafa"

	serious := true 

	fmt.Printf("%s will win in %d! Does he believe: %t. Expected rate: %g...", name, year, serious, expectedRate)

	// Mustafa will win in 2026! Does he believe: true. Expected rate: 99.99...
}
