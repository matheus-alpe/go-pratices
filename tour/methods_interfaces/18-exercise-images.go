package methodsinterfaces

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h    int
	r, g, b uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.r, i.g, i.b, 255}
}

func ExerciseImagesExample() {
	fmt.Println("\nExercise Images:")

	m := Image{w: 100, h: 100, r: 200, g: 150, b: 40}
	pic.ShowImage(m)
}
