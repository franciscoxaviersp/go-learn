package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{a, b}

	for i, v1 := range x {
		for j, v2 := range v1 {
			fmt.Println(i, j, v2)
		}
	}
}
