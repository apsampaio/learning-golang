package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is a test err")

	r := err.Error()

	fmt.Println(r)
}
