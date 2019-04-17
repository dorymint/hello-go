// Pointer receivers.
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
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) NotAffect() {
	v.X = 0
	v.Y = 0
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // be changed.
	fmt.Println(v.Abs())

	v.NotAffect() // not changed.
	fmt.Println(v.Abs())
}
