package methodsinterfaces

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func StringersExample() {
	fmt.Println("\nStringers:")

	a := Person{"Matheus Alves", 27}
	z := Person{"Thiago Laurenir", 21}
	fmt.Println(a)
	fmt.Println(z)
}
