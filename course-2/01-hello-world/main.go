package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	Foo()
	Bar()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func Foo() {
	fmt.Println("Im in foo")
}

func Bar() {
	fmt.Println("Im in bar")
}
