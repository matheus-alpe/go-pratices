package basics

import "fmt"

func PointersExample() {
	fmt.Println("\nPointers:")

	i, j := 42, 2701

	p := &i
	fmt.Println(*p, p)
	*p = 21
	fmt.Println(i)

	fmt.Println(j)
	p = &j
	*p = *p / 27
	fmt.Println(j)
}
