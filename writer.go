package main

import "fmt"

func main() {
	var v consoleWriter = consoleWriter{}
	v.Write([]byte("Hello, Go"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) Write(data []byte) (int, error) {

	n, error := fmt.Println(string(data))
	return n, error
}
