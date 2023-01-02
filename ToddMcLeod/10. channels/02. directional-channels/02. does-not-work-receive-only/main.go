package main

import "fmt"

func main() {

	// Here we create a send only type channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// here we are trying to receive from the channel
	// which causes the error ❌
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
