package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	counter := 0

	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	// define the mutex variable
	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			// locks the counter variable
			mu.Lock()
			v := counter
			// tell the program to go to another task if has one
			runtime.Gosched()
			v++
			counter = v
			// unlock the counter variable
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
