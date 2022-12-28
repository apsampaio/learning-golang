package main

import "fmt"

var x func()

func main() {

	x := func() {
		fmt.Println("Hello from X")
	}

	x()

}
