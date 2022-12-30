package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Println("Types:")
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a

	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored in the address
	fmt.Println(*&a)

	*b = 43 // get value from address b and assign new value

	fmt.Println(a)
}
