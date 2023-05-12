package main

import (
	"fmt"
)

func main() {
	// anonymous func
	c1 := make(chan int)

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)

	//buffered channel
	c2 := make(chan int)

	go func() {
		c2 <- 42
	}()

	fmt.Println(<-c2)
}
