package main

import (
	"fmt"
	"sync"
)

func main() {

	c := 0
	const gr = 10
	var wg sync.WaitGroup

	wg.Add(gr)

	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			//
			mu.Lock()

			v := c
			v++
			c = v
			fmt.Println(c)

			//
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", c)

}
