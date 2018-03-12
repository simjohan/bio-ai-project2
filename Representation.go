package main

import "fmt"

type NewIndividual struct {
	SegmentMatrix   [][]int
	DirectionMatrix [][]Direction
}

func initNewIndividual(segments [][]Vertex, directions map[Vertex]Direction) NewIndividual {
	return NewIndividual{generateSegmentMatrix(segments), nil}
}

func generateSegmentMatrix(segments [][]Vertex) [][]int {
	segmentMatrix := make([][]int, pictureWidth)

	for i, segment := range segments {
		segmentMatrix = append(segmentMatrix, make([]int, pictureHeight))
		for _, node := range segment {
			fmt.Print(node.X,node.Y)
			segmentMatrix[node.X][node.Y] = i
		}
	}

	return segmentMatrix
}