package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42
	// channels is blocking here and cannot continue

	fmt.Println(<-c)
}
