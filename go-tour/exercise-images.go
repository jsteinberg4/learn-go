// TODO: Exercise: Images (https://go.dev/tour/methods/25)
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	// TODO: Implement interface method
	return nil
}

func (img Image) Bounds() image.Rectangle {
	// TODO: Implement interface method
	return image.Rect(0, 0, 0, 0)
}

func (img Image) At(x, y int) color.Color {
	// TODO: Implement interface method
	return nil
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
