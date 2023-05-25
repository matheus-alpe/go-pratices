package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 27
	return &p
}

func main() {
	p1 := newPerson("John")
	fmt.Println(p1)
	fmt.Println(p1.age)

	fmt.Println(person{"Sophia", 45})
	fmt.Println(person{name: "Mark", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s)

	pointerP1 := &s
	fmt.Println(pointerP1)
	pointerP1.age = 25
	fmt.Println(pointerP1)
}
