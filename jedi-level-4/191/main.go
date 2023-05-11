package main

import "fmt"

func main() {

	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("1: ")
	print(x)

	fmt.Println("2: ")
	x["petronilho_francisco"] = []string{"f1", "f2", "f3"}

	print(x)

	fmt.Println("3: ")
	delete(x, "no_dr")

	print(x)

}

func print(dict map[string][]string) {
	for i, v1 := range dict {
		fmt.Printf("%v: ", i)
		for j, v2 := range v1 {
			fmt.Printf("%v, %v\t", j, v2)
		}
		fmt.Println("")
	}
}
