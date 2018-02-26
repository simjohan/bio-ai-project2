package main

import (
	"math/rand"
	"fmt"
	"math"
	"image"
)

type Direction int

/* can be replaced with Point{} */
type Point struct {
	x, y int
}

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

// Returns a Point-array containing the cardinal neighbour points
func getNeighbourhood(point Point, height, width int) []Point {
	var adjacent []Point

	// up
	if point.x != 0 {
		adjacent = append(adjacent, Point{point.x - 1, point.y})
	}

	// down
	if point.x != height-1 {
		adjacent = append(adjacent, Point{point.x + 1, point.y})
	}

	// left
	if point.y != 0 {
		adjacent = append(adjacent, Point{point.x, point.y - 1})
	}

	// right
	if point.y != width-1 {
		adjacent = append(adjacent, Point{point.x, point.y + 1})
	}

	return adjacent
}


func minKey(from Point, neighbours []Point, picture *Picture) Point {
	// set initially to infinity
	min := math.Inf(0)
	var leastCostNeighbour Point

	for i := range neighbours {
		origin := &picture.pixels[from.x][from.y]
		to := &picture.pixels[neighbours[i].x][neighbours[i].y]
		rgbDistance := euclideanDistance(origin, to)
		//fmt.Println(rgbDistance)
		//fmt.Println(rgbDistance)
		if rgbDistance < min {
			min = rgbDistance
			leastCostNeighbour = neighbours[i]
		}
	}
	//fmt.Println(leastCostNeighbour)
	return leastCostNeighbour
}

// Implement Prims algorithm to make a minimum spanning tree
func generateMinimumSpanningTree(picture *Picture) {

	var parent [2][2]Point
	var cost    [2][2]float64
	var visited [2][2]bool
/*
	var parent [321][481]Point
	var key    [321][481]float64
	var mstSet [321][481]bool
*/
	// init all keys as infinite
	for row := range cost {
		for col := range cost[row] {
			cost[row][col] = math.Inf(0)
			visited[row][col] = false
		}
	}

	cost[0][0] = 0
	parent[0][0] = Point{}
	visited[0][0] = true

	/*
	        mstSet[u] = true;

        // Update key value and parent index of the adjacent vertices of
        // the picked vertex. Consider only those vertices which are not yet
        // included in MST
        for (int v = 0; v < V; v++)

           // graph[u][v] is non zero only for adjacent vertices of m
           // mstSet[v] is false for vertices not yet included in MST
           // Update the key only if graph[u][v] is smaller than key[v]
          if (graph[u][v] && mstSet[v] == false && graph[u][v] <  key[v])
             parent[v]  = u, key[v] = graph[u][v];
     }

     // print the constructed MST
     printMST(parent, V, graph);
}
	 */
	 for row := 0; row < len(picture.pixels); row++ {
		for col := 0; col < len(picture.pixels[0]); col++ {
			pos := Point{col, row}
			neighbours := getNeighbourhood(pos, picture.width, picture.height)
			fmt.Println(picture.pixels[col][row])
			u := minKey(pos, neighbours, picture)
			//visited[u.x][u.y] = true
			fmt.Println(neighbours)
			for n := 0; n < len(neighbours)-1; n++ {
				dist := euclideanDistance(&picture.pixels[col][row], &picture.pixels[neighbours[n].y][neighbours[n].x])
				//fmt.Println(dist, neighbours[n].x, neighbours[n].y)
				if visited[neighbours[n].y][neighbours[n].x] == false &&
					dist < cost[neighbours[n].y][neighbours[n].x] {
						visited[neighbours[n].y][neighbours[n].x] = true
						parent[neighbours[n].y][neighbours[n].x] = u
						cost[neighbours[n].y][neighbours[n].x] = dist
						//fmt.Println(key[neighbours[n].x][neighbours[n].y])
				}
			}

		}
	}
	fmt.Println(parent)
	fmt.Println(cost)
	fmt.Println(visited)

}
