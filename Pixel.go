package main

import (
	"math"

	"image/color"
)

type Pixel struct {
	R, G, B, A uint8
}

type Picture struct {
	pixels        [][]Pixel
	width, height int
	numPixels     int
}

func rgbaToPixel(r, g, b, a uint8) Pixel {
	return Pixel{r, g, b, a}
}

/*func euclideanDistance(p1, p2 *Pixel) float64 {
	return math.Sqrt(math.Pow(float64(p1.R-p2.R), 2) + math.Pow(float64(p1.G-p2.G), 2) + math.Pow(float64(p1.B-p2.B), 2))
}*/

func euclideanDistance(p1 *Pixel, p2 *Pixel) float64 {
	r := math.Pow(float64(p1.R)-float64(p2.R), 2)
	g := math.Pow(float64(p1.G)-float64(p2.G), 2)
	b := math.Pow(float64(p1.B)-float64(p2.B), 2)
	a := math.Pow(float64(p1.A)-float64(p2.A), 2)
	return math.Sqrt(r + g + b + a)
}

func updatePixelColors(segments [][]Vertex, picture Picture) {
	for segment := range segments {
		color := averageSegmentColor(segments[segment], &picture)
		r,g,b,a := color.RGBA()

		for element := range segments[segment] {
			x := segments[segment][element].X
			y := segments[segment][element].Y
			picture.pixels[x][y] = Pixel{uint8(r),uint8(g),uint8(b),uint8(a)}

		}
	}
}

func updateEdgeColors(segment []Vertex, picture Picture) {
	for element := range segment {
		x := segment[element].X
		y := segment[element].Y
		picture.pixels[x][y] = Pixel{uint8(229),uint8(241),uint8(0),uint8(255)}

	}
}

func averageSegmentColor(segment []Vertex, picture *Picture) color.Color {
	var r, g, b int

	for point := range segment {
		x := segment[point].X
		y := segment[point].Y
		pixel := picture.pixels[x][y]

		r += int(pixel.R)
		g += int(pixel.G)
		b += int(pixel.B)
	}

	r /= len(segment)
	g /= len(segment)
	b /= len(segment)


	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
