package main

import "fmt"

type person struct {
	first_name    string
	last_name     string
	fav_ice_cream []string
}

func main() {

	persons := map[string]person{}
	joao := person{
		first_name:    "joao",
		last_name:     "antonio",
		fav_ice_cream: []string{"morango", "chocolate", "vanilla"},
	}

	luis := person{
		first_name:    "luis",
		last_name:     "eduardo",
		fav_ice_cream: []string{"ovos moles", "laranja", "avela"},
	}

	persons[joao.last_name] = joao
	persons[luis.last_name] = luis
	print(persons)
}

func print(dict map[string]person) {
	for i, v1 := range dict {
		fmt.Printf("%v favourite flavors: ", i)
		for _, v2 := range v1.fav_ice_cream {
			fmt.Printf("%v, ", v2)
		}
		fmt.Println("")

	}
}
