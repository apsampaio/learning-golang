package main

import "fmt"

// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavor

type person struct {
	first_name                 string
	last_name                  string
	favorite_ice_cream_flavors []string
}

func main() {

	p := person{
		first_name:                 "Andre",
		last_name:                  "Sampaio",
		favorite_ice_cream_flavors: []string{"Chocolate", "Strawberry"},
	}

	for _, v := range p.favorite_ice_cream_flavors {
		fmt.Println(v)
	}

}
