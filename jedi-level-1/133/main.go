package main

import "fmt"

type coldcat int

var x coldcat
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println(int(x))
	fmt.Printf("%T\n", int(x))
	y = int(x)
	fmt.Println(y)
}
