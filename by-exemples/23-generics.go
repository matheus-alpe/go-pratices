package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}

	return elements
}

func main() {
	ints := map[string]int64{
		"first":  23,
		"second": 12,
	}

	fmt.Println("SumInts:", SumInts(ints))

	floats := map[string]float64{
		"first":  23.23,
		"second": 12.11,
	}

	fmt.Println("SumFloats:", SumFloats(floats))

	fmt.Println("SumIntsOrFloats (ints):", SumIntsOrFloats(ints))
	fmt.Println("SumIntsOrFloats (floats):", SumIntsOrFloats(floats))
	fmt.Println()

	var m = map[int]string{
		1: "2",
		2: "4",
		4: "8",
	}

	fmt.Println("keys:", MapKeys(m))
	// MapKeys[int, string](m)

	fmt.Println()

	// singly-linked list
	fmt.Println("singly-linked list")

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(32)
	fmt.Println("list:", list.GetAll())

	list1 := List[string]{}
	list1.Push("Matheus")
	list1.Push("Marcos")
	list1.Push("Thiago")
	fmt.Println("list1:", list1.GetAll())
}
