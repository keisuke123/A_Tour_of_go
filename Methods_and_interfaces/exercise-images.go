package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Point struct {
	X, Y int
}

type Image struct {
	Min, Max Point
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.Min.X, i.Min.Y, i.Max.X, i.Max.Y)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) / 2), 255, 255, 255}
}

func main() {
	m := Image{Point{0, 0}, Point{500, 500}}
	pic.ShowImage(m)
}
