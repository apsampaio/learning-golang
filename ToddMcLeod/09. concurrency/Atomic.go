package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	var counter int64

	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
