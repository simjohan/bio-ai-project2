package main

import (
	"math/rand"
	"fmt"
	"math"
)

type Direction int

type Position struct {
	x, y int
}

type Gene struct {
	position  Position
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

// Returns a Position-array containing the cardinal neighbour positions
func getNeighbourhood(position Position, height, width int) []Position {
	var adjacent []Position

	// up
	if position.x != 0 {
		adjacent = append(adjacent, Position{position.x - 1, position.y})
	}

	// down
	if position.x != height-1 {
		adjacent = append(adjacent, Position{position.x + 1, position.y})
	}

	// left
	if position.y != 0 {
		adjacent = append(adjacent, Position{position.x, position.y - 1})
	}

	// right
	if position.y != width-1 {
		adjacent = append(adjacent, Position{position.x, position.y + 1})
	}

	return adjacent

	/*for row := x; row < x; row++ {
	 		for col := y; col < y; row++ {

			}
		}*/
}

func minKey(from Position, neighbours []Position, picture *Picture) Position {
	// set initially to infinity
	min := math.Inf(0)
	var leastCostNeighbour Position

	for i := range neighbours {
		origin := &picture.pixels[from.x][from.y]
		to := &picture.pixels[neighbours[i].x][neighbours[i].y]
		rgbDistance := euclideanDistance(origin, to)
		//fmt.Println(rgbDistance)
		if rgbDistance < min {
			min = rgbDistance
			leastCostNeighbour = neighbours[i]
		}
	}
	return leastCostNeighbour
}

// Implement Prims algorithm to make a minimum spanning tree
func generateMinimumSpanningTree(picture *Picture) {
	//randRow := rand.Intn(picture.height)
	//randCol := rand.Intn(picture.width)

	fmt.Println(picture.width)

	var parent [10][10]int
	var key    [10][10]float64
	var mstSet [10][10]bool

	// init all keys as infinite
	for row := range key {
		for col := range key[row] {
			key[row][col] = math.Inf(0)
			mstSet[row][col] = false
		}
	}

	key[0][0] = 0
	parent[0][0] = -1

	for row := range picture.pixels {
		for col := range picture.pixels[row] {
			pos := Position{row, col}
			neighbours := getNeighbourhood(pos, picture.height, picture.width)
			u := minKey(pos, neighbours, picture)
			mstSet[u.x][u.y] = true

		}
	}
	for row := range picture.pixels {
		for col := range picture.pixels[row] {

			fmt.Println(mstSet[row][col])

		}
	}
}
