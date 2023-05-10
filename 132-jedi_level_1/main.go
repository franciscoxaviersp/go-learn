package main

import "fmt"

type coldcat int

var x coldcat

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
