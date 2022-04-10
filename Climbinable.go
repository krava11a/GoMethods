package main

import (
	"fmt"
	"math"
)

type Climbinable interface {
	Abs() float64
}

type IceAxe float64

func (ia IceAxe) Abs() float64 {
	if ia < 0 {
		return float64(-ia)
	}
	return float64(ia)

}

type Hill struct {
	X, Y float64
}

func (h *Hill) Abs() float64{
	return math.Sqrt(h.X*h.X + h.Y*h.Y)
}


func TestInterfaces() {
	var c Climbinable
	ia := IceAxe(-math.Sqrt2)
	h := Hill{3,4}

	c = ia
	fmt.Println("ia",c.Abs())
	c = &h

	fmt.Println("hill",c.Abs())

}
