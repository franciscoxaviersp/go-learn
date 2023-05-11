package main

import "fmt"

func main() {
	p1 := struct {
		first_name string
		last_name  string
	}{
		first_name: "joao",
		last_name:  "luis",
	}

	fmt.Println(p1)
}
