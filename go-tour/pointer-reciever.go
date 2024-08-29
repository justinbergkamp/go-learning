package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (test Vertex) Abs() float64 {
	return math.Sqrt(test.X*test.X + test.Y*test.Y)
}

func (test *Vertex) Scale(f float64) {
	test.X = test.X * f
	test.Y = test.Y * f
}

func main() {
	test := Vertex{3, 4}
	test.Scale(10)
	fmt.Println(test.Abs())
}
