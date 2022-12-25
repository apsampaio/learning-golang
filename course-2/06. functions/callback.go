package main

import "fmt"

func main() {
	ii := []int{1, 2, 3}
	s := sum(ii...)

	fmt.Println(s)

	s2 := even(sum, ii...)

	fmt.Println(s2)
}

func sum(xi ...int) int {
	t := 0

	for _, v := range xi {
		t += v
	}

	return t
}

func even(cb func(xi ...int) int, vi ...int) int {

	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return cb(yi...)
}
