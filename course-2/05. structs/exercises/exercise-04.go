package main

import "fmt"

//Create and use an anonymous struct.

func main() {

	p := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "Andre",
		last_name:  "Sampaio",
		age:        25,
	}

	fmt.Println(p)

}
