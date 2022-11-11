package main

import "fmt"

func main() {

	// slice allows you to group together VALUES of the same TYPE
	s := []int{1, 2, 3} // composite literal

	// using a for with a slice

	for i, v := range s {
		fmt.Printf("Index: %d | Value: %d\n", i, v)
	}

	// accessing a specific index

	fmt.Print("The value of index 0 is 1? R: ")
	fmt.Println(s[0] == 1)

	// append to slice

	s = append(s, 4)

	fmt.Println("new value appended to s: ", s)

	x := []int{5, 6, 7}

	// appending a slice to another using "spread operator"

	s = append(s, x...)
	fmt.Println("new slice x: ", x)

	fmt.Println("appeding x to s: ", s)

	// removing data from slice from 0 to 3

	s = append(s[0:3])

	fmt.Println("removing from s and returning to 3 values: ", s)

	// we can also create a predefined slice using the make command

	y := make([]int, 3, 3)

	fmt.Println(len(y))
	fmt.Println(cap(y))

	// we can also have multi-dimensional slices

	coord_x := []int{30, 45}
	coord_y := []int{10, 10}

	coord := [][]int{coord_x, coord_y}

	fmt.Println(coord)
}
