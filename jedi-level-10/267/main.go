package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 2
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	go func() {
		c <- 3
	}()

	v, ok = <-c
	fmt.Println(v, ok)
}
