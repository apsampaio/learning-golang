package main

import "fmt"

// We define a interface that demands a functions called greet
// for each struct created we define a greet method
// when the structs corresponds to the methods of the interface
// it fits without any declaration

type person struct {
	name string
}

type dog struct {
	name string
}

type alive interface {
	greet()
}

func (d dog) greet() {
	fmt.Println("BARK!")
}

func (p person) greet() {
	fmt.Println("Hi, my name is", p.name)
}

func sayHello(a alive) {
	a.greet()
}

func main() {
	p := person{
		name: "Andre",
	}

	d := dog{
		name: "Rex",
	}

	sayHello(p)
	sayHello(d)
}
