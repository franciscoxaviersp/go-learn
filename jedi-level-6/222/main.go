package main

import "fmt"

func main() {

	f := incrementor()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
