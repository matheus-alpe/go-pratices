package methodsinterfaces

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// https://go.dev/tour/methods/22
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // rune
	}

	return len(b), nil
}

func ExerciseReadersExample() {
	fmt.Println("\nExercise Readers:")

	reader.Validate(MyReader{})
}
