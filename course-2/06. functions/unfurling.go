package main

import "fmt"

func main() {

	xi := []int{1, 2, 3}

	// foo needs 0 or more ints not a int slice
	// sum := foo(xi) ❌
	sum := foo(xi...) // ✅

	fmt.Println(sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}
