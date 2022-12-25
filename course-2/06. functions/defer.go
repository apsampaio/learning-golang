package main

import "fmt"

func main() {
	// creates a definition where the program should end
	// when we run this code, foo is going to run after bar
	defer foo()
	bar()
}

func foo() {
	fmt.Println("FOO")
}

func bar() {
	fmt.Println("BAR")
}
