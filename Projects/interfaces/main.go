package main

import "fmt"

// Interfaces are used to define a method set
// Every type with the same methods of a interfaces, becomes the interface
type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Imaginary custom logic
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// Imaginary custom logic
	return "Hola"
}
