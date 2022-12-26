package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("HELLO FROM FOO")
}

func bar() {
	fmt.Println("HELLO FROM BAR")

}
