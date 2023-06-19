package generics

import (
	"fmt"
)

type List[T any] struct {
	val  T
	next *List[T]
}

func (l *List[T]) Len() int {
	count := 0

	// iterate through all nodes
	for node := l; node != nil; node = node.next {
		count++
	}

	return count
}

func (l *List[T]) InsertAt(pos int, value T) *List[T] {
	if l == nil || pos <= 0 {
		return &List[T]{
			val:  value,
			next: l,
		}
	}

	l.next = l.next.InsertAt(pos-1, value)
	return l
}

func (l *List[T]) Append(value T) *List[T] {
	return l.InsertAt(l.Len(), value)
}

func (l *List[T]) String() string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%v->%v", l.val, l.next.String())
}

func GenericTypesExample() {
	fmt.Println("\nGeneric types:")

	var stringList = &List[string]{val: "Matheus"}
	stringList.Append("A")
	fmt.Println(stringList)
	stringList.Append("B")
	fmt.Println(stringList)

	var intList = &List[int]{val: 1}
	intList.Append(2)
	fmt.Println(intList)
	intList.Append(3)
	fmt.Println(intList)
}
