package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  23,
		"second": 12,
	}

	fmt.Println(ints)

	m := make(map[string]int)
	fmt.Println("m", m)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	m["k3"] = 27

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs2:", prs)

	_, prs = m["k1"]
	fmt.Println("prs1:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
