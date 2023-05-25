package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r1 := rect{width: 20, height: 10}

	fmt.Println("area:", r1.area())
	fmt.Println("perimeter:", r1.perimeter())

	rp := &r1
	rp.width = 50
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
