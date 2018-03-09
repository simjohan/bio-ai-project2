package main
/*
import (
	"errors"
	"fmt"
	"sort"
)

type Node struct {
	x, y int
}

type Edge struct {
	from, to Node
	cost float64
}


func getNeighbours(from Node) (Node, error, Node, error) {
	var downError, rightError error
	var down, right Node

	// down
	if from.x != pictureWidth-1 {
		down = Node{from.x + 1, from.y}
	} else {
		downError = errors.New("no downward neighbour")
	}
	// right
	if from.y != pictureHeight-1 {
		right = Node{from.x, from.y + 1}
	} else {
		rightError = errors.New("no right neighbour")
	}
	return right, rightError, down, downError

}


func makeMst(picture *Picture) ([]Edge) {

	// nodes (pixels), the bool is to keep track of nodes visited
	nodes := make(map[Node]bool)
	// all existing edges
	edges := make([]Edge,0)
	graph := make([]Edge, 0)

	// make edges
	for x := 0; x < pictureWidth; x++ {
		for y := 0; y < pictureHeight; y++ {
			node := Node{x, y}
			nodes[node] = false

			right, rightError, down, downError := getNeighbours(node)
			if rightError == nil {
				distance := euclideanDistance(&picture.pixels[node.x][node.y], &picture.pixels[right.x][right.y])
				edges = append(edges, Edge{node, right, distance})
			}
			if downError == nil {
				distance := euclideanDistance(&picture.pixels[node.x][node.y], &picture.pixels[down.x][down.y])
				edges = append(edges, Edge{node, down, distance})
			}
		}
	}

	/* Kruskal-ish
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

 	for len(graph) < len(nodes)-1 {

		selected := edges[0]
		edges = edges[1:]

		if !nodes[selected.to] {
			nodes[selected.to] = true
			//inGraph[selected] = true
			graph = append(graph, selected)
		}

	}

	fmt.Println(len(nodes), len(graph))

	return graph

}*/

