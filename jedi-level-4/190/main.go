package main

import "fmt"

func main() {

	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, v1 := range x {
		fmt.Printf("%v: ", i)
		for j, v2 := range v1 {
			fmt.Printf("%v, %v\t", j, v2)
		}
		fmt.Println("")
	}

	x["petronilho_francisco"] = []string{"f1", "f2", "f3"}

	for i, v1 := range x {
		fmt.Printf("%v: ", i)
		for j, v2 := range v1 {
			fmt.Printf("%v, %v\t", j, v2)
		}
		fmt.Println("")
	}
}
