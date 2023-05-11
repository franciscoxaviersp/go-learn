package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(list ...int) int {

	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum

}

func bar(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
