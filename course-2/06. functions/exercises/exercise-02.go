package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}

func foo(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}

func bar(xi []int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
