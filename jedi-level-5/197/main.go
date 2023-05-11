package main

import "fmt"

type person struct {
	first_name    string
	last_name     string
	fav_ice_cream []string
}

func main() {

	joao := person{
		first_name:    "joao",
		last_name:     "antonio",
		fav_ice_cream: []string{"morango", "chocolate", "vanilla"},
	}

	fmt.Println(joao.first_name, joao.last_name)
	fmt.Print("Favourite flavors: ")
	for _, v := range joao.fav_ice_cream {
		fmt.Printf("%v, ", v)
	}
	fmt.Println("")

	luis := person{
		first_name:    "luis",
		last_name:     "eduardo",
		fav_ice_cream: []string{"ovos moles", "laranja", "avela"},
	}

	fmt.Println(luis.first_name, luis.last_name)
	fmt.Print("Favourite flavors: ")
	for _, v := range luis.fav_ice_cream {
		fmt.Printf("%v, ", v)
	}
	fmt.Println("")

}
