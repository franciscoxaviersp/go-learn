package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	count := 100
	wg.Add(count)

	incrementor := 0

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			temp := incrementor

			incrementor = temp + 1
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
