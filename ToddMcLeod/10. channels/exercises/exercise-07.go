package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
