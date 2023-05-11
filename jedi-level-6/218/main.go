package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Testing anonumous func", x)
	}(43)
}
