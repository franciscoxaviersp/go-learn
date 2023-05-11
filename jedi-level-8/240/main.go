package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type UsersByAgeLast []user

func (u UsersByAgeLast) Len() int {
	return len(u)
}
func (u UsersByAgeLast) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByAgeLast) Less(i, j int) bool {
	if u[i].Age != u[j].Age {
		return u[i].Age < u[j].Age
	}
	return u[i].Last < u[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "M",
		Last:  "AMMMM",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	u3 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u4 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3, u4}

	fmt.Println(users)

	for _, user := range users {
		fmt.Printf("User: %v %v, aged %v\n", user.First, user.Last, user.Age)
		fmt.Println("\t", user.Sayings)
	}

	// your code goes here
	sort.Sort(UsersByAgeLast(users))
	fmt.Println("")
	fmt.Println("Sorted users by age and last")

	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Printf("User: %v %v, aged %v\n", user.First, user.Last, user.Age)
		fmt.Println("\t", user.Sayings)
	}
}
