package methodsinterfaces

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

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsExample() {
	fmt.Println("\nMethods:")

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}
