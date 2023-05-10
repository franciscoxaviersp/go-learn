package main

import "fmt"

const initial_year = 2023 - 4

const (
	year1 = iota + initial_year
	year2
	year3
	year4
)

func main() {
	fmt.Println(year1, year2, year3, year4)
}
