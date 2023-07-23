// TODO: Exercise: Images (https://go.dev/tour/methods/25)
// Instructions:
// > Remember the picture generator you wrote earlier? Let's write another one,
//
//	but this time it will return an implementation of image.Image instead of a slice of data.
//
// > Define your own Image type, implement the necessary methods, and call pic.ShowImage.
//
//	See: https://go.dev/pkg/image/#Image for "necessary methods"
//
// > `Bounds` should return a `image.Rectangle`, like `image.Rect(0, 0, w, h)`.
// > `ColorModel` should return `color.RGBAModel`.
// > `At` should return a color; the value `v` in the last picture generator corresponds to
//
//	`color.RGBA{v, v, 255, 255}` in this one.
package exercises

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	// TODO: Implement interface method
	return image.Rect(0, 0, 320, 240)
}

func (img Image) At(x, y int) color.Color {
	// TODO: Implement interface method
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func ColorImage() {
	m := Image{}
	pic.ShowImage(m)
}
