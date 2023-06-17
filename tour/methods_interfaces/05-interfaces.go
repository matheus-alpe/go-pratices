package methodsinterfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex1 struct {
	x, y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func InterfacesExample() {
	fmt.Println("\nInterfaces:")

	var a Abser

	f := MyFloat2(-math.Sqrt2)
	v := Vertex1{2, 3}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
}
