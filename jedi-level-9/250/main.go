package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	count := 100
	wg.Add(count)

	incrementor := 0

	for i := 0; i < 100; i++ {
		go func() {
			temp := incrementor
			runtime.Gosched()
			incrementor = temp + 1
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
