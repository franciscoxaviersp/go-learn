package main

import "fmt"

func main() {

	f := foo()

	fmt.Println(f(23, 25))
}

func foo() func(x ...int) int {

	return func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}
		return sum
	}

}
