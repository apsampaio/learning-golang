package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct {
}

func main() {
	arg := os.Args[1]

	lw := logWriter{}
	f, _ := os.Open(arg)

	io.Copy(lw, f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
