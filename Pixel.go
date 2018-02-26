package main

import "math"

type Pixel struct {
	R, G, B, A uint8
}

type Picture struct {
	pixels        [][]Pixel
	width, height int
	numPixels    int
}

func rgbaToPixel(r, g, b, a uint8) Pixel {
	return Pixel{r, g, b, a}
}

/*func overallDeviation(segmentSet [][]int) {

	for k := 0; k < len(segmentSet); k++ {
		for i := 0; i < len(segmentSet[k]); i++ {

		}
	}

}
*/

func euclideanDistance(p1, p2 *Pixel) float64 {
	return math.Sqrt(math.Pow(float64(p1.R-p2.R), 2) + math.Pow(float64(p1.G-p2.G), 2) + math.Pow(float64(p1.B-p2.B), 2))
}
