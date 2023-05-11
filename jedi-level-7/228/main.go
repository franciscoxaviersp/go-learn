package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "joao",
		last:  "antonio",
	}

	fmt.Println(&p)
	fmt.Printf("%T\n", &p)

	fmt.Println(p)

	changeMe(&p, "luis")

	fmt.Println(p)
}

func changeMe(pp *person, new_first string) {
	pp.first = new_first
}
