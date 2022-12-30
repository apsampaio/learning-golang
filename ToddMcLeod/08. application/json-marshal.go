package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Andre",
		Last:  "Sampaio",
		Age:   25,
	}

	p2 := person{
		First: "Eliete",
		Last:  "Sampaio",
		Age:   49,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
