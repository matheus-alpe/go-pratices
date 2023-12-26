package tags

import "fmt"

type Person struct {
	Name string `example:"name"`
}

func (u *Person) String() string {
	return fmt.Sprintf("Hi! My name is %v", u.Name)
}

func StructTagsExample() {
	u := &Person{"Matt"}
	fmt.Println(u)
}
