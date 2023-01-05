package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	c := 0
	const gr = 10
	var wg sync.WaitGroup

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := c
			runtime.Gosched()
			v++
			c = v
			fmt.Println(c)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", c)

}
