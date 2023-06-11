package basics

import "fmt"

type Location struct {
	Lat, Long float64
}

var locations = map[string]Location{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		Lat: 37.42202, Long: -122.08408,
	},
}

var m map[string]Location

func exampleMap1() {
	fmt.Println("\nExample 1:")

	fmt.Println(locations)

	m = make(map[string]Location)
	m["Bell labs"] = Location{
		40.68433, -74.39967,
	}

	fmt.Println(m)
	fmt.Println(m["Bell labs"])

	m1 := make(map[string]Location)
	fmt.Println(m1)
}

func exampleMap2() {
	fmt.Println("\nExample 2:")

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func MapExample() {
	fmt.Println("\nMap:")

	exampleMap1()
	exampleMap2()
}
