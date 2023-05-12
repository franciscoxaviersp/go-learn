package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func (p *Person) speak() {
	fmt.Println("Hello, I'm", p.First)
}

type Human interface {
	speak()
}

func main() {
	p := Person{
		First: "Luis",
		Last:  "Antonio",
	}

	saySomething(&p)

	// does not work
	// saySomething(p)
}

func saySomething(h Human) {
	h.speak()
}
