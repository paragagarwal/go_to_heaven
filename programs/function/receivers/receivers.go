package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Sum(v *Vertex) (f float64) {
	return v.X+v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Sum(&v))
	v.Scale(10)
	fmt.Println(Sum(&v))
	fmt.Println(v.Abs())
}
