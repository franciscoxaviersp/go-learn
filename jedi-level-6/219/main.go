package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("Testing anonumous func", x)
	}
	f(32)
}
