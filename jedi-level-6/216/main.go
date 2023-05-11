package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and I'm %v years old\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "luis",
		last:  "eduardo",
		age:   43,
	}
	p1.speak()
}
