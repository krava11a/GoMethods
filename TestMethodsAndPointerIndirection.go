package main

import (
	"fmt"
	"math"
)

type Mountain struct {
	X,Y float64
}

func (m *Mountain) ScaleMountain(f float64)  {
	m.X = m.X * f
	m.Y = m.Y * f
}

func ScaleFuncMountain(m *Mountain, f float64){
	m.X = m.X * f
	m.Y = m.Y * f
}

func (m *Mountain) AbsMountain() float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

func AbsFuncMountain(m Mountain) float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}



func TestMethodsAndPointerIndirection() {
	fmt.Println("________________________TestMethodsAndPointerIndirection")
	m := Mountain{3, 4}
	//m.ScaleMountain(2)
	//ScaleFuncMountain(&m, 10)


	//p := &Mountain{4, 3}
	//p.ScaleMountain(3)
	//ScaleFuncMountain(p, 8)
	//fmt.Println(m, p)
	//
	//fmt.Println(m.AbsMountain())
	//fmt.Println(AbsFuncMountain(m))
	//
	//fmt.Println(p.AbsMountain())
	//fmt.Println(AbsFuncMountain(*p))


	fmt.Printf("Before scaling: %+v, Abs: %v\n", m, m.AbsMountain())
	m.ScaleMountain(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", m, m.AbsMountain())

}