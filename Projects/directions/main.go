package main

import "fmt"

func main() {
	var s [11]int

	for i := range s {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
