package main
import "fmt";

func main() {

	greetVoid()

	greetPerson("Aykan")

	sum(6, 12, 40)
}

func greetVoid() {

	fmt.Println("Hi there!")
}

func greetPerson(name string) {

	fmt.Printf("Hi %s!\n", name)
}

func sum(x int, y int, z int) {

	fmt.Printf("%d + %d + %d = %d\n", x, y, z, x + y + z)
}

// Hi there!
// Hi Aykan!
// 6 + 12 + 40 = 58
