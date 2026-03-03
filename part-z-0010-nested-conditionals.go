package main
import (
	"strings"
	"fmt"
);

func main() {

	city := "Antalya"
	age := 19
	idle := false

	if strings.ToLower(city) == "ankara" {

		if age >= 18 && idle == true {

			fmt.Println("We can hire")

		} else {

			fmt.Println("We can't hire...")
		}

	} else {

		fmt.Println("We are in Ankara! We can't hire him/her...")
	}
}
