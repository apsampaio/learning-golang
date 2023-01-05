package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42
	// as we only reserved 1 value for the channel
	// the code gets blocked after you try to assign a second value
	c <- 43

	fmt.Println(<-c)
}
