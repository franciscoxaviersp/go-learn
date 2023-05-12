package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	go foo()
	fmt.Println(runtime.NumGoroutine())
	go bar()
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	fmt.Println("Foo is running")
	wg.Done()
}

func bar() {
	fmt.Println("Bar is running")
	wg.Done()
}
