package main

import "fmt"

// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice

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

	p2 := person{
		first_name:                 "John",
		last_name:                  "Doe",
		favorite_ice_cream_flavors: []string{"Mint", "Strawberry"},
	}

	m := map[string]person{
		p.last_name:  p,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v)
	}

}
