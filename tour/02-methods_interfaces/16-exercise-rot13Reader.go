package methodsinterfaces

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err != nil {
		return n, err
	}

	for i, value := range b {
		if value >= 'N' && value <= 'Z' || value >= 'n' && value <= 'z' {
			b[i] -= 13
			continue
		}

		if value >= 'A' && value <= 'M' || value >= 'a' && value <= 'm' {
			b[i] += 13
		}
	}

	return n, nil
}

func ExerciseRot13ReaderExample() {
	fmt.Println("\nExercise rot13Reader:")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
