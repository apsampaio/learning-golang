package main

import "fmt"

func main() {
	foo(bar)
}

func foo(x func() int) {
	fmt.Println(x())
}

func bar() int {
	return 42
}
