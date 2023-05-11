package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * 2
}

type shape interface {
	area() float64
}

func info(s shape) {

	switch s := s.(type) {
	case circle:
		fmt.Println("Circle area: ", s.area())
	case square:
		fmt.Println("Square ares: ", s.area())
	}

}

func main() {
	square := square{
		length: 1.5,
		width:  2.3,
	}

	circle := circle{
		radius: 3,
	}

	fmt.Println(square.area())
	fmt.Println(circle.area())

	info(square)
	info(circle)
}
