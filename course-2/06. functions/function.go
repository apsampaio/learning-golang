package main

import "fmt"

func main() {
	print("Hi mom")

	x, y := foo("Hi")
	fmt.Println(x, y)
}

// function (r receiver) identifier(parameters) (returns(s)) {...}

// everything in go is PASS BY VALUE
func print(s string) bool {
	fmt.Println(s)

	return true
}

func foo(s string) (bool, string) {
	return true, s
}
