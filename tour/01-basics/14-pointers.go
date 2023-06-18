package basics

import "fmt"

/*
>> The * operator is used to declare a pointer type or to dereference a pointer.
|--> Declaration: When defining a variable, * is used to denote that the variable is a pointer to a certain type. For example, var p *int declares a variable p of type "pointer to int".
|--> Dereferencing: When a pointer is declared, the * operator can be used to access the value it points to. For example, if p is a pointer to an int, *p gives the value stored at the memory address pointed to by p.

>> The & operator is used to obtain the memory address of a variable.
|--> Addressing: The & operator, when placed before a variable, returns the memory address where the variable is stored. For example, if x is a variable, &x gives the memory address of x.
*/

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
