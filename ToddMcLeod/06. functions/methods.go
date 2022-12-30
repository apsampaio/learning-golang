package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person person
	ltk    bool
}

// func (r receiver) identifier(parameters) (returns(r)) { code}

// here we attach the method speak to the struct secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa1.speak()
}
