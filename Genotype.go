package main

import (
	"math/rand"
	"fmt"
)

type Direction int

type Gene struct {
	x, y      int
	direction Direction
}

type Chromosome struct {
	genes []Gene
}

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

func generateRandomDirection() Direction {
	switch rand.Intn(5) {
	case 0:
		return Up
	case 1:
		return Down
	case 2:
		return Left
	case 3:
		return Right
	case 4:
		return None
	}
	return None
}

/*func generateRandomGenotype(image [][]Pixel) Chromosome {
	height := len(image)
	width := len(image[0])
	genes := make([]Gene, height*width)

	i := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			genes[i] = Gene{col, row, generateRandomDirection()}
			i++
		}
	}

	return Chromosome{genes}
}*/

/*
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[x][y] = rgbaToPixel(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
 */

/*func makeWeightTable(pixels [][]Pixel) [][]float64 {
	height := len(pixels)
	width := len(pixels[0])
	weightTable := make([][]float64, width)

	for i := 0; i < width; i++ {
		weightTable[i] = make([]float64, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			weightTable[x][y] = 323.2
		}
	}
	fmt.Println("")

	return weightTable
}*/

// Implement Prims algorithm to make a minimum spanning tree
func generateMinimumSpanningTree(image *[][]Pixel) {

}
