package main

import "fmt"

func main() {
	// we specify a buffer channel to allow one value to not block the channel
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
