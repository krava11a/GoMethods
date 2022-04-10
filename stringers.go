package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestStringers() {
	a:= Person{"Arti Dent",42}
	z:= Person{"Zapha Beeble",9002}
	fmt.Println(a,z)
}
