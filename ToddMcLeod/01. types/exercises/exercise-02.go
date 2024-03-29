package main

import "fmt"

/*

Use var to DECLARE three VARIABLES. The variables should have package level
scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
variables and make sure the variables are of the following TYPE (meaning they can
store VALUES of that TYPE).
	a. identifier “x” type int
	b. identifier “y” type string
	c. identifier “z” type bool
2. in func main
	a. print out the values for each identifier
	b. The compiler assigned values to the variables. What are these values called?

*/

var x int
var y string
var z bool

func main() {

	fmt.Println("Zero values: ")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
