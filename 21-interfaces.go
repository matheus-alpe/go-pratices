package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
	fmt.Println()
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (d *Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (d Llama) Speak() string {
	return "??!"
}

type Programmer struct{}

func (d Programmer) Speak() string {
	return "Design Patterns!"
}

// func DoSomething(v interface{}) {
// 	// ... https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
// }

func PrintAll(values []interface{}) {
	for _, val := range values {
		fmt.Println(val)
	}
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}

	*t = Timestamp(v)
	return nil
}

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	r := rectangle{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)

	animals := []Animal{new(Dog), new(Cat), Llama{}, Programmer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	fmt.Println()

	names := []string{"stanley", "david", "oscar"}
	values := make([]interface{}, len(names))
	for i, v := range names {
		values[i] = v
	}
	PrintAll(values)
	fmt.Println()

	var input = `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012"
	}
	`

	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
	fmt.Println()

	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s.String())
}
