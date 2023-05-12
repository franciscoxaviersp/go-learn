package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	// does not run
	// fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	cr := make(<-chan int)

	go func() {
		// does not run
		// cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
