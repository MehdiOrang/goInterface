package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incremnter = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}

type IntCounter int

type Incremnter interface {
	Increment() int
}

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
