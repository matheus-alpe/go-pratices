package basics

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p1 = &Vertex{1, 2}
)

func StructExample() {
	fmt.Println("\nStructs:")

	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 4
	fmt.Println(v.x)

	p := &v
	p.x = 1e9
	fmt.Println(v)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p1)
}
