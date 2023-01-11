package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
}

func mySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
