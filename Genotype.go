package main
/*
import (
	"fmt"
	"math"
	"errors"
	"log"
)

/* can be replaced with Point{}
type Point struct {
	x, y int
}


// Returns a Point-array containing the cardinal neighbour points
func getNeighbourhood(point Point, width, height int) ([]Point){
	var adjacent []Point

	// down
	if point.x != height-1 {
		adjacent = append(adjacent, Point{point.x + 1, point.y})
	}

	// left
	if point.x != 0 {
		adjacent = append(adjacent, Point{point.x - 1, point.y})
	}

	// right
	if point.y != width-1 {
		adjacent = append(adjacent, Point{point.x, point.y + 1})
	}

	// up
	if point.y != 0 {
		adjacent = append(adjacent, Point{point.x, point.y - 1})
	}

	return adjacent
}

func minKey(from Point, mstSet map[Point]bool, picture *Picture) (Point, float64, error) {
	// set initially to infinity
	min := math.Inf(0)
	var leastCostNeighbour Point

	neighbours := getNeighbourhood(from, picture.width, picture.height)
	if len(neighbours) == 0 {
		fmt.Println("neighbours was empty", from)
		return Point{}, min, errors.New("no minkey (end of graph)")
	}


	for i := range neighbours {
		if mstSet[neighbours[i]] == false {
			origin := &picture.pixels[from.x][from.y]
			to := &picture.pixels[neighbours[i].x][neighbours[i].y]
			rgbDistance := euclideanDistance(origin, to)
			if rgbDistance < min {
				min = rgbDistance
				leastCostNeighbour = neighbours[i]
			}
		}
	}

	return leastCostNeighbour, min, nil
}


func makeMatrix(picture *Picture, width, height int) [481][321]float64 {
	var matrix [481][321]float64
	for x := 0; x < width-1; x++ {

		for y := 0; y < height-1; y++ {
			//log.Println(x, y)
			adjacent := getNeighbourhood(Point{x, y}, width, height)
			for a := range adjacent {
				//log.Println(a)
				matrix[x][y] = euclideanDistance(&picture.pixels[x][y], &picture.pixels[adjacent[a].x][adjacent[a].y])
			}
		}
	}

	for x := range matrix {
		fmt.Println(matrix[x])
	}
	return matrix
}

// Implement Prims algorithm to make a minimum spanning tree
func generateMinimumSpanningTree(picture *Picture, width, height int) { // graph *[][]float64

	var parent = make(map[Point]Point)
	var key = make(map[Point]float64)
	var mstSet = make(map[Point]bool)
	//graph := makeMatrix(picture, width, height)

	// init all keys as inf
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := Point{x, y}
			key[p] = math.Inf(0)
			mstSet[p] = false
		}
	}

	// use the first pixel as root node
	/*node := Point{0,0}
	key[node] = 0
	parent[node] = node


	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			origin := Point{x, y}
			mstSet[origin] = true

			u, _, err := minKey(origin, mstSet, picture)
			if err != nil {
				log.Println(err)
			}

			adjacent := getNeighbourhood(u, width, height)

			for a := 0; a < len(adjacent); a++ {
				// get the point (key) of current neighbour
				v := adjacent[a]
				distance := euclideanDistance(&picture.pixels[u.x][u.y], &picture.pixels[v.x][v.y])
				p := Point{0, 1}
				if v == p {
					fmt.Println(key[v])
				}
				if mstSet[v] == false && distance < key[v] {
					parent[v] = u
					key[v] = distance
				} else {
					if mstSet[v] == false {
						//fmt.Println(v, "FALSE")
					}
				}
			}
		}

	}
	/*for k, v := range key {
		fmt.Println("k: ", k, "v", v)
	}
	for kid, parent := range parent {
		fmt.Println(parent, "->", kid)
	}



}*/


