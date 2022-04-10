package main

import (
	"fmt"
	"math"
)

func Abs(vertex Vertex) float64 {
	return math.Sqrt(vertex.X*vertex.X + vertex.Y*vertex.Y)
}

func Scale(vertex *Vertex, f float64) {
	vertex.X = vertex.X * f
	vertex.Y = vertex.Y * f
}

func TestPointersAndFunctions() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}


