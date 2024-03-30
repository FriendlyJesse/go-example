package main

import (
	"fmt"
)

type Foo interface {
	Say()
}

type Dog struct {
	name string
}

func (d Dog) Say() {
	fmt.Println(d.name + " say hi")
}

func main() {

	// example.Sampling()
	// example.Loterry()
	// example.Calc()
}
