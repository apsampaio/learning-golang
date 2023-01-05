package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from function one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from function two")
		wg.Done()
	}()

	wg.Wait()
}
