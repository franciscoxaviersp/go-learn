package main

import "fmt"

type customErr struct {
	custom_string string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("this is a custom error: %v\n", ce.custom_string)
}

func main() {
	ce := customErr{
		custom_string: "This is an error",
	}

	foo(ce)
}

func foo(err error) {
	fmt.Printf("Inside foo: %v", err.Error())
	fmt.Println("Custom string:", err.(customErr).custom_string)
}
