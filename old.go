package main

import (

)
/*
type Pair struct {
	Key Point
	Value float64
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

type Gene struct {
	point  image.Point
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

/*
int minKey(int key[], bool mstSet[])
{
   // Initialize min value
   int min = INT_MAX, min_index;

   for (int v = 0; v < V; v++)
     if (mstSet[v] == false && key[v] < min)
         min = key[v], min_index = v;

   return min_index;
}
 *//*
func minK(key map[Point]float64, mstSet map[Point]bool) Point {
	min := math.Inf(0)
	var minIndex Point
	for k, _ := range key {
		if mstSet[k] == false && key[k] < min {
			min = key[k]
			minIndex = k
		}
	}
	return minIndex
}

func min(cost map[Point]float64, mstSet map[Point]bool) PairList {
	pl := make(PairList, len(cost))
	i := 0
	for k, v := range cost {
		if mstSet[k] == false {
			pl[i] = Pair{k, v}
		}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
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
}


func makeMatrix(picture *Picture, width, height int) [3][3]float64 {
	var matrix [3][3]float64
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			adjacent := getNeighbourhood(Point{x, y}, width, height)
			for a := range adjacent {
				matrix[x][y] = euclideanDistance(&picture.pixels[x][y], &picture.pixels[adjacent[a].x][adjacent[a].y])
			}
		}
	}
	fmt.Println(matrix)
	return matrix
}

*/

/*
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[x][y] = rgbaToPixel(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}
*/
