package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Hi mom")
	fmt.Fprintln(os.Stdout, "Hi mom")
	io.WriteString(os.Stdout, "Hi mom")

}
