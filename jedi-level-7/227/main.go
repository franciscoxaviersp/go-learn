package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 43
	fmt.Println(x)
}
