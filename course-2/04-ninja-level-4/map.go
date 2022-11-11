package main

import "fmt"

func main() {

	// creating a map

	m := map[string]int{
		"Andre": 24,
		"John":  25,
	}

	fmt.Println(m)

	// accessing a map key by name
	fmt.Println(m["Andre"])

	// checking if a map key exists
	v, ok := m["Doe"]
	fmt.Println(v, ok)

	// checing if a map key existing one liner
	if v, ok := m["John"]; ok {
		fmt.Println("THIS IS JOHN's AGE: ", v)
	}

	// adding data to map
	m["Doe"] = 26

	v, ok = m["Doe"]

	fmt.Println(v, ok)

	// using map with a for loop

	for i, v := range m {
		fmt.Println(i, v)
	}

	// deleting a map key

	delete(m, "Doe")

	v, ok = m["Doe"]

	fmt.Println(v, ok)

}
