package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func TestInterfacesAreSatisfiedImplicity() {
	var i I = T{"Axer"}
	i.M()
}
