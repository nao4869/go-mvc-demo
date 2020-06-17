package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	// once we lanuch, execute this
	go func(input chan string) {
		input <- "send to channel"
	}(c)

	// main go routine - listen the result
	test := <-c
	fmt.Println(test)
}

// HelloWorld -
func HelloWorld() {
	fmt.Println("Hello World")
}
