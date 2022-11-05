package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values
const (
	year = 2022
)
const (
	_ = iota
	a = year + iota
	b = year + iota
	c = year + iota
	d = year + iota
)

func main() {
	fmt.Println(a, b, c, d)

}
