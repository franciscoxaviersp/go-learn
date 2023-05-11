package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	foo(func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}
		return sum
	}, x...)
}

func foo(f func(...int) int, x ...int) {
	fmt.Println("Sum is:", f(x...))
}
