package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello, my name is ", p.first, p.last, ". I'm ", p.age, "years old.")
}

func main() {
	p := person{
		first: "Andre",
		last:  "Sampaio",
		age:   25,
	}

	p.speak()
}
