package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

// Abs() float64を実装している
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// vはVertex(*Vertexではない)なので, Abserを実装してない
	// a = v
	fmt.Println(a.Abs())
}
