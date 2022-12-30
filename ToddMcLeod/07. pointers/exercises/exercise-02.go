package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		name: "Andre",
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.name = "Name Changed"
	// (*p).name = "Name Changed" // this is also valid ✅
}
