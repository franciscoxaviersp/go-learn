package main

import "fmt"

func main() {
	x := 42

	if x == 40 {
		fmt.Println("x equals 40")
	} else if x == 41 {
		fmt.Println("x equals 41")
	} else {
		fmt.Println("x is different from 40 and 41")
	}
}
