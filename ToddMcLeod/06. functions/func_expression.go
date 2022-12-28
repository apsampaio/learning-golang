package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("I'm a function assigned to a variable")
	}

	x()
}
