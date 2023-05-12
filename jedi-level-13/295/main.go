package main

import (
	"fmt"

	"github.com/franciscoxaviersp/go-learn/jedi-level-13/295/quote"
	"github.com/franciscoxaviersp/go-learn/jedi-level-13/295/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
