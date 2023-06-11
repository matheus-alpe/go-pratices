package basics

import (
	"fmt"
	// "golang.org/x/tour/pic"
)

// https://go.dev/tour/moretypes/18
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for x := range s {
		s[x] = make([]uint8, dx)

		for y := range s[x] {
			s[x][y] = uint8((x * y) * (x * y) / 2)
		}
	}

	return s
}

func PrintExerciseSlices() {
	fmt.Println("\nExercise 02: Slices")

	// pic.Show(Pic)
}
