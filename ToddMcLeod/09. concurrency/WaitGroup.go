package main

import (
	"fmt"
	"runtime"
	"sync"
)

// receivers 	values
// (t T)		T and *T
// (t *T) 		*T

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	// ask to waitgroup to wait for 1 thing
	wg.Add(1)

	go foo()
	bar()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	// wait for waitgroup to be done
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	// tell that waitgroup is done
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
