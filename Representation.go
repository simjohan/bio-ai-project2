package main

import (
	"math/rand"
)

type MatrixIndividual struct {
	SegmentMatrix    [][]int
	DirectionMatrix  [][]Direction
	SegmentMap		 map[int][]Vertex
	overallDeviation float64
	edgeValue 		 float64
	Rank 			 int
	CrowdingDistance float64
	Fitness			 float64
	Dominates		 []*MatrixIndividual
	DominatedBy 	 int

}

func (i *MatrixIndividual) Init(graph Graph) {
	randomKValue := rand.Intn((6000 - 5000) + 5000)
	_, i.DirectionMatrix = graph.GraphSegmentation(randomKValue)
	i.SegmentMatrix, i.SegmentMap = DirectionMatrixToSegmentMatrixAndSegmentMap(i.DirectionMatrix)
	i.CalculateFitness()
}

func (i *MatrixIndividual) IsDominating(i2 *MatrixIndividual) bool {
	if i.edgeValue == i2.edgeValue && i.overallDeviation == i2.overallDeviation {
		return false
	}
	return i.edgeValue >= i2.edgeValue && i.overallDeviation <= i2.overallDeviation
}

func (i *MatrixIndividual) mutate() {

	// Choose a random segment to merge

	seg1ID := randomSegmentId(i.SegmentMap)
	seg2ID := seg1ID

	// Find the second one
	visited := make(map[Vertex]bool)
	opened := make(map[Vertex]bool)
	queue := []Vertex{i.SegmentMap[seg1ID][rand.Intn(len(i.SegmentMap[seg1ID]))]}
	for len(queue) > 0 && seg2ID == seg1ID {
		n := queue[0]
		visited[n] = true
		queue = queue[1:]
		neighbours := getAllCardinalNeighbours(n)
		for neighbour := range neighbours {
			if _, visited := visited[neighbours[neighbour]]; visited {
				continue
			}
			if _, o := opened[neighbours[neighbour]]; !o {
				queue = append(queue, neighbours[neighbour])
				opened[neighbours[neighbour]] = true
				segID := i.SegmentMatrix[neighbours[neighbour].X][neighbours[neighbour].Y]
				if segID != seg1ID {
					seg2ID = segID
					continue
				}
			}
		}
	}
	if seg1ID == seg2ID {
		return
	}
	// Merge them
	for _, n := range i.SegmentMap[seg2ID] {
		i.SegmentMatrix[n.X][n.Y] = seg1ID
	}
	i.SegmentMap[seg1ID] = append(i.SegmentMap[seg1ID], i.SegmentMap[seg2ID]...)
	delete(i.SegmentMap, seg2ID)
}


func (i *MatrixIndividual) CalculateFitness() {
	i.edgeValue = i.EdgeValue()
	i.overallDeviation = i.OverallDeviation()
	i.Fitness = (i.edgeValue * float64(edgeWeight)) + (i.overallDeviation * float64(deviationWeight))
}

func (i *MatrixIndividual) OverallDeviation() float64 {
	deviation := 0.0
	for _, segment := range i.SegmentMap {
		centeroid := averageSegmentColor(segment, pic)
		p2 := rgbaToPixel(centeroid)
		for _, node := range segment {
			p1 := pic.pixels[node.X][node.Y]
			deviation += euclideanDistance(&p1, &p2)
		}
	}
	return deviation / float64(len(i.SegmentMap))
}

func (i *MatrixIndividual) EdgeValue() float64 {
	edgeValue := 0.0

	for segmentId := range i.SegmentMap {
		for _, node := range i.SegmentMap[segmentId] {
			p1 := pic.pixels[node.X][node.Y]
			neighbours := getAllCardinalNeighbours(node)
			for _, neighbour := range neighbours {
				if inBounds(neighbour) {
					if i.SegmentMatrix[node.X][node.Y] != i.SegmentMatrix[neighbour.X][neighbour.Y] {
						p2 := pic.pixels[neighbour.X][neighbour.Y]
						edgeValue += euclideanDistance(&p1, &p2)
					}
				}
			}

		}
	}

	return edgeValue / float64(len(i.SegmentMap))
}


func randomSegmentId(segmentMap map[int][]Vertex) int {
	seg := rand.Intn(len(segmentMap))
	counter := 0
	for k := range segmentMap {
		if counter == seg {
			seg = k
		}
		counter++
	}
	return seg
}

//
/* just for converting back to the old repr for drawing */
func (i MatrixIndividual) segMapToDraw() [][]Vertex {
	result := make([][]Vertex,0)
	for seg := range i.SegmentMap {
		result = append(result, i.SegmentMap[seg])
	}
	return result
}

func (i MatrixIndividual) segMapToIdMapDraw() [][]Vertex {
	result := make([][]Vertex,0)
	for seg := range i.SegmentMap {
		result = append(result, i.SegmentMap[seg])
	}
	return result
}
