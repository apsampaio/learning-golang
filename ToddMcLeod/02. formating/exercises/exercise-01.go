package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex
// fmt print formatting https://pkg.go.dev/fmt

func main() {

	x := 42

	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %#x\n", x)

}
