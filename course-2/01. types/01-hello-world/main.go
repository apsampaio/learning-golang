package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	Foo()
	Bar()
}

// using var we can declare outside of a function
var y = 24

// declaring a variable within a type without assigning a value
// by default it gives a 'zero value' to this variable
// each type has a different 'zero value' default
// such as 0 for Int or 0.0 for float
var z int

func Foo() {
	// Declaring a variable and assigning at the same time
	// using short declaration we can only declare inside the function
	x := 42
	fmt.Println(x, y, z)
}

func Bar() {
	var a = "Hi"
	// Go is a STATIC PROGRAMMING LANGUAGE
	//  it's all about TYPES
	// here we can check the type of a variable
	fmt.Printf("%T\n", a)
}
