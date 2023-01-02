package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("about to exit")
}

// send
// here we are creating a scope where the channel is a send only
func foo(c chan<- int) {
	c <- 42
}

// receive
// here we are creating a scope where the channel is a receive only
func bar(c <-chan int) {
	fmt.Println(<-c)
}
