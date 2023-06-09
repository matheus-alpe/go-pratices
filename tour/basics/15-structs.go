package basics

import "fmt"

type Vertex struct {
	x, y int
}

func StructExample() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 4
	fmt.Println(v.x)

	p := &v
	p.x = 1e9
	fmt.Println(v)
}
