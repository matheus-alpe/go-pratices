package methodsinterfaces

import (
	"fmt"
	"math"
)

/*
-- There are two reasons to use a pointer receiver.

1: is so that the method can modify the value that its receiver points to.

2: is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
*/

func abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func absPointer(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func absFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PointersFunctionsExample() {
	fmt.Println("\nPointers Functions:")

	v := Vertex{2, 4}
	scale(&v, 10)
	fmt.Println(abs(v))
	fmt.Println(v.Abs())

	v.Scale(2) // = (&v).Scale(2)
	fmt.Println(v)
	(&v).Scale(2)
	fmt.Println(v)
	scale(&v, 2)
	fmt.Println(v)
	fmt.Println("absFunc", absFunc(v))

	p := &Vertex{4, 3}
	p.Scale(2)
	fmt.Println(p)
	scale(p, 2)
	fmt.Println(p)
	fmt.Println("absFunc", absFunc(*p)) // Dereferencing and accessing his value

	fmt.Println(v, p)

	g := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", g, absPointer(g))
	g.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", g, absPointer(g))
}
