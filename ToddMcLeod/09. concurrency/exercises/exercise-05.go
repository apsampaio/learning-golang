package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var c int64
	const gr = 10

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&c, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&c))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter: ", c)

}
