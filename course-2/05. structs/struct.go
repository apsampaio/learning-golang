package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	job
}

type job struct {
	profile string
	salary  int
}

func main() {

	p := person{
		first_name: "Andre",
		last_name:  "Sampaio",
		job: job{
			profile: "Programmer",
			salary:  3000,
		},
	}

	// anonimous struct
	dog := struct {
		name  string
		breed string
	}{
		name:  "Bob",
		breed: "Husky",
	}

	fmt.Println(p)
	fmt.Println(dog)

}
