package main

import (
	"fmt"
	"time"
)

func main() {

	go hello()
	go hello()
	go hello()
	go hello()
	go hello()
	go hi()
	go hi()
	go hi()
	
	time.Sleep(5 * time.Second)
}

func hello() {
	fmt.Println("Hello")
}

func hi() {
	fmt.Println("Hi")
}

/*
Hello
Hello
Hello
Hello
Hi   
Hi   
Hello
Hi

*/


