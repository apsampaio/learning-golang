package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("I: ", i)
	}

	x := 0

	// that represents a while loop
	for x < 10 {
		x++
		fmt.Println(x)
	}
}
