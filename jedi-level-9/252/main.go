package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	count := 100
	wg.Add(count)

	var incrementor int64

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			runtime.Gosched()
			fmt.Printf("Counter: %v\n", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
