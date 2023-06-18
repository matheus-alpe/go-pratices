package methodsinterfaces

import "fmt"

type I1 interface {
	M()
}

type T1 struct {
	S string
}

/*
This method means type T implements the interface I,
but we don't need to explicitly declare that it does so.
*/
func (t *T1) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

type F1 float64

func (f F1) M() {
	fmt.Println(f)
}

func InterfacesNilExample() {
	fmt.Println("\nInterfaces Nil:")

	var i2 I1
	var t *T1
	i2 = t
	describe(i2)
	i2.M()

	i2 = &T1{"world"}
	describe(i2)
	i2.M()

	/*
		A nil interface value holds neither value nor concrete type.

		Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	*/
	// var i I1
	// describe1(i)
	// i.M()
}
