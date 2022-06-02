package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {

	p := person{
		firstName: "Andre",
		lastName:  "Sampaio",
		contact: contact{
			email:   "andre@sampaio.com",
			zipCode: 94000,
		},
	}

	p.updateName("Jim")
	//p.print()

	name := "Bill"
	fmt.Println(*&name)
}

func (pPointer *person) updateName(new string) {
	(*pPointer).firstName = new
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
