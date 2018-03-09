package main

import (
	"errors"

	"fmt"
)

func getNeighbours(from Point) (Point, error, Point, error) {
	var downError, rightError error
	var down, right Point

	// right
	if from.X != pictureWidth-1 {
		right = Point{from.X + 1, from.Y}
	} else {
		rightError = errors.New("no right neighbour")
	}
	// down
	if from.Y != pictureHeight-1 {
		down = Point{from.X, from.Y + 1}
	} else {
		downError = errors.New("no down neighbour")
	}
	return right, rightError, down, downError

}

func makeGraph(picture *Picture) Graph {

	var edges    []Edge
	var vertices []Vertex

	for y := 0; y < pictureHeight; y++ {
	for x := 0; x < pictureWidth; x++ {

			from := Point{x, y}
			vertices = append(vertices, from)
			right, rightError, down, downError := getNeighbours(from)
			if rightError == nil {
				distance := euclideanDistance(&picture.pixels[from.X][from.Y], &picture.pixels[right.X][right.Y])
				edges = append(edges, Edge{from, right, distance})
			}/* else {
				log.Println(rightError, from.X, from.Y, picture.pixels[from.X][from.Y])
			}*/
			if downError == nil {
				distance := euclideanDistance(&picture.pixels[from.X][from.Y], &picture.pixels[down.X][down.Y])
				edges = append(edges, Edge{from, down, distance})
			} /*else {
				log.Println(downError, from.X, from.Y, picture.pixels[from.X][from.Y])
			}*/
		}
	}
	return Graph{edges, vertices}
}



// global vars
var pictureWidth, pictureHeight int

func main() {
	//picture, _ := readImageFromFile("test3x3.jpg")
	picture, _ := readImageFromFile("2/Test image.jpg")
	pictureWidth = picture.width
	pictureHeight = picture.height


	graph := makeGraph(&picture)
	fmt.Println(graph.Edges)

	segments := graph.GraphSegmentation(1500)
	fmt.Println(averageSegmentColor(segments[0], &picture))


	updatePixelColors(segments, picture)
	drawPicture(picture)

/*	for x := range segs {
		for e := 0; e < x-1; e++ {
			fmt.Println(segs[x][e].(Point).X)
		}
	}*/

	/*log.Println("Started...")
	start := time.Now()

	graph.GraphSegmentation(2)


	elapsed := time.Since(start)
	log.Printf("execution time %s", elapsed)


	//blue := Pixel{0, 54, 255, 255}
	//green := Pixel{0, 255, 48, 255}
	//red := Pixel{255, 0, 0, 255}
	//fmt.Println("from blue to red:", euclideanDistance(&blue, &red))
	//fmt.Println("from green to red:", euclideanDistance(&green, &red))
	//generateMinimumSpanningTree(&picture, picture.width, picture.height)



	/*log.Println("started")
	//makeMatrix(&picture, picture.width, picture.height)
	graph := makeMst(&picture)
	log.Println("ended")
	//drawPicture(&picture)
	cost := 0.0
	for r := range graph {
		//fmt.Println(graph[r].cost)
		cost += graph[r].cost
	}
	fmt.Println(cost)*/

}
