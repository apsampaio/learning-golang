package main

import "fmt"

func main() {
	c := make(chan int)

	// here we create another go routine
	// it doesnt block the main rotine and the codes continue runing
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
