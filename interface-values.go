package main

import (
	"fmt"
	"math"
)

type i interface {
	Me()
	NotMe()
}

type t struct {
	st string
}

func (t *t) Me() {
	fmt.Println(t.st)
}
func (t *t) NotMe() {
	fmt.Println(t.st)
}

type F float64

func (f F) Me() {
	fmt.Println(f)
}
func (f F) NotMe() {
	fmt.Println(f)
}


func TestInterfaceValues() {
	var interf i
	interf = &t{"Sema"}
	describe(interf)
	interf.Me()

	interf = F(math.Pi)
	describe(interf)
	interf.Me()
}

func describe(i i) {
	fmt.Printf("(%v, %T)\n",i,i)
}
