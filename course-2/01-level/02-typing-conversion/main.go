package main

import "fmt"

// creating our own type
type potato int

var a int
var b potato

func main() {

	b = 1
	a = 2

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// we can't directly assign 'b' to 'a' because both have different types
	// but we can convert the type 'b' to match the 'a' type

	// a = b 	❌
	a = int(b) // ✅
}
