package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			for j := 0; j < 10; j++ {
				c <- j*10 + v
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
