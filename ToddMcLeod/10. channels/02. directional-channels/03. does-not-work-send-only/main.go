package main

import "fmt"

func main() {

	// Here we create a receive only type channel
	c := make(<-chan int, 2)

	// here we are trying to send to the channel
	// which causes the error ❌
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
