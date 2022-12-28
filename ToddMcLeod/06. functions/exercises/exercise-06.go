package main

import "fmt"

func main() {

	fmt.Println(func(x int) int {
		return x
	}(42))

}
