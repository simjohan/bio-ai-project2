package main

import (
	"sort"
	"math"
	"fmt"
	"errors"
)

type Edge struct {
	U, V   Vertex
	Weight float64
}

//type Vertex = interface{}
type Vertex struct {
	X, Y int
}

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}

type Point struct {
	X, Y int
}

type Border struct {
	ColMin, RowMin, ColMax, RowMax map[int]int
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f\n", e.U, e.V, e.Weight)
}

func (g Graph) GraphSegmentation(k int) ([][]Direction, [][]int, map[int][]Vertex) {
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*Element)
	maxInternal := make(map[*Element]float64)
	sizeMap     := make(map[*Element]int)

	// this is the genotype, i.e flat list of directions
	mst := make([]Edge, 0)

	for _, vertex := range g.Vertices {
		element := MakeSet(vertex)
		sizeMap[element] = 1
		disjointMap[vertex] = element
	}

	for _, edge := range g.Edges {
		from, to := FindSet(disjointMap[edge.U]), FindSet(disjointMap[edge.V])
		if from != to {
			minInt := math.Min(maxInternal[from] + float64(k/sizeMap[from]), maxInternal[to] + float64(k/sizeMap[to]))
			if edge.Weight  <= minInt {
				mst = append(mst, edge)
				size1, size2 := sizeMap[from], sizeMap[to]
				Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
			}
		}
	}

	/* stitch together segments made from k-kruskal */
	segMap := make(map[*Element][]Vertex)
	for _,vertex := range g.Vertices {
		rep := FindSet(disjointMap[vertex])
		segMap[rep] = append(segMap[rep], vertex)

	}

	segments := make([][]Vertex,0)
	for _, value := range segMap {
		segments = append(segments, value)
	}


	/* start on building flat map (genotype)  */
	directions 	   := make(map[Vertex]Direction)
	directionsFlat := make([]Direction, 0)
	parent         := make(map[Vertex]Vertex)

	/* init all node parents to itself */
	for _, node := range g.Vertices {
		parent[node] = node
	}

	neighbours := make(map[Vertex][]Vertex)
	for _, edge := range mst {
		neighbours[edge.U]  = append(neighbours[edge.U], edge.V)
		neighbours[edge.V] = append(neighbours[edge.V], edge.U)
	}

	/* bfs to make a tree with all vertices, with only one edge */
	for _, segment := range segments {
		V := make(map[Vertex]bool)
		Q := []Vertex{segment[0]}
		/* Init queue with the first node in order to spawn the loop */
		for len(Q) > 0 {
			node := Q[0]
			V[node] = true
			Q = Q[1:]
			for _, neighbour := range neighbours[node] {
				if _, V := V[neighbour]; !V {
					Q = append(Q, neighbour)
					parent[neighbour] = node
				}
			}
		}
	}

	for node, par := range parent {
		directions[node] = edgeDirection(Edge{node, par, 0})
		directionsFlat = append(directionsFlat, edgeDirection(Edge{node, par, 0}))
	}


	segmentMatrix := make([][]int, pictureWidth)
	for x := range segmentMatrix {
		segmentMatrix[x] = make([]int, pictureHeight)
	}

	directionMatrix := make([][]Direction, pictureWidth)

	i := 0
	for x := 0; x < pictureWidth; x++ {
		directionMatrix[x] = make([]Direction, pictureHeight)
		for y := 0; y < pictureHeight; y++ {
			//fmt.Println(i)
			directionMatrix[x][y] = directionsFlat[i]
			i++
		}
	}

	segmendIdMap := make(map[int][]Vertex)
	for x, segment := range segments {
		segmendIdMap[x] = segment
		for _, elem := range segment {
			segmentMatrix[elem.X][elem.Y] = x
		}
	}


	return directionMatrix, segmentMatrix, segmendIdMap
}

func edgeDirection(edge Edge) Direction {
	if edge.U.X < edge.V.X {
		return Right
	}
	if edge.U.X > edge.V.X {
		return Left
	}
	if edge.U.Y < edge.V.Y {
		return Down
	}
	if edge.U.Y > edge.V.Y {
		return Up
	}
	return None
}


type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

func genoToPheno(directions []Direction) [][]Vertex {

	directionsMap := make(map[Vertex]Direction)
	i := 0
	for x := 0; x < pictureWidth; x++ {
		for y := 0; y < pictureHeight; y++ {
			//fmt.Println(i)
			directionsMap[Vertex{x, y}] = directions[i]
			i++
		}
	}

	disjoint := make(map[Vertex]*Element)
	for v := range directionsMap {
		simpleSet := MakeSet(disjoint[v])
		disjoint[v] = simpleSet
	}

	for v, dir := range directionsMap {
		adjacent := getNeighbourFromDirection(v, dir)
		if inBounds(adjacent) {
			set1 := FindSet(disjoint[v])
			set2 := FindSet(disjoint[adjacent])
			Union(set1, set2)
		}
	}

	toSegment := make(map[*Element][]Vertex)
	for s := range directionsMap {
		toSegment[FindSet(disjoint[s])] = append(toSegment[FindSet(disjoint[s])], s)
	}

	segments := make([][]Vertex, 0)
	for _, v := range toSegment {
		segments = append(segments, v)
	}
	//segmentIdMap := generateSegmentIdMap(segments)

	return segments

}

func getNeighbours(from Vertex) (Vertex, error, Vertex, error) {
	var downError, rightError error
	var down, right Vertex

	// right
	if from.X != pictureWidth-1 {
		right = Vertex{from.X + 1, from.Y}
	} else {
		rightError = errors.New("no right neighbour")
	}
	// down
	if from.Y != pictureHeight-1 {
		down = Vertex{from.X, from.Y + 1}
	} else {
		downError = errors.New("no down neighbour")
	}
	return right, rightError, down, downError

}


// TODO: consider only returning bottom and right neighbours
func getAllCardinalNeighbours(from Vertex) []Vertex {
	nodes := make([]Vertex, 0)

	// right
	if from.X != pictureWidth-1 {
		nodes = append(nodes, Vertex{from.X + 1, from.Y})
	}

	// left
	if from.X > 0 {
		nodes = append(nodes, Vertex{from.X - 1, from.Y})
	}

	// down
	if from.Y != pictureHeight-1 {
		nodes = append(nodes, Vertex{from.X, from.Y + 1})
	}

	// up
	if from.Y > 0 {
		nodes = append(nodes, Vertex{from.X, from.Y - 1})
	}

	return nodes

}

func getTwoCardinalNeighbours(from Vertex) []Vertex {
	nodes := make([]Vertex, 0)

	// right
	if from.X != pictureWidth-1 {
		nodes = append(nodes, Vertex{from.X + 1, from.Y})
	}

	// left
	//if from.X > 0 {
	//	nodes = append(nodes, Vertex{from.X - 1, from.Y})
	//}

	// down
	if from.Y != pictureHeight-1 {
		nodes = append(nodes, Vertex{from.X, from.Y + 1})
	}

	// up
	//if from.Y > 0 {
	//	nodes = append(nodes, Vertex{from.X, from.Y - 1})
	//}

	return nodes

}

func makeGraph(picture *Picture) Graph {

	var edges    []Edge
	var vertices []Vertex

	for y := 0; y < pictureHeight; y++ {
		for x := 0; x < pictureWidth; x++ {

			from := Vertex{x, y}
			vertices = append(vertices, from)
			right, rightError, down, downError := getNeighbours(from)
			if rightError == nil {
				distance := euclideanDistance(&picture.pixels[from.X][from.Y], &picture.pixels[right.X][right.Y])
				edges = append(edges, Edge{from, right, distance})
			}
			if downError == nil {
				distance := euclideanDistance(&picture.pixels[from.X][from.Y], &picture.pixels[down.X][down.Y])
				edges = append(edges, Edge{from, down, distance})
			}
		}
	}
	return Graph{edges, vertices}
}

func getNeighbourFromDirection(node Vertex, direction Direction) Vertex {
	switch direction {
	case Up:
		return Vertex{node.X, node.Y - 1}
	case Down:
		return Vertex{node.X, node.Y + 1}
	case Left:
		return Vertex{node.X - 1, node.Y}
	case Right:
		return Vertex{node.X + 1, node.Y}
	default:
		return node
	}
}



func SegmentMatrixAndSegmentMapToDirectionMatrix(segmentMatrix [][]int, segmentMap map[int][]Vertex) [][]Direction {
	matrix := make([][]Direction, pictureWidth)
	for i := range matrix {
		matrix[i] = make([]Direction, pictureHeight)
	}
	for x := 0; x < pictureWidth; x++ {
		for y := 0; y < pictureHeight; y++ {
			matrix[x][y] = None
		}
	}
	for id := range segmentMap {
		visited := make(map[Vertex]bool)
		opened := make(map[Vertex]bool)
		queue := []Vertex{segmentMap[id][0]}
		for len(queue) > 0 {
			n := queue[0]
			visited[n] = true
			queue = queue[1:]
			neighbours := getAllCardinalNeighbours(n)
			for i := range neighbours {
				if _, visited := visited[neighbours[i]]; visited {
					continue
				}
				if _, o := opened[neighbours[i]]; !o {
					if segmentMatrix[neighbours[i].X][neighbours[i].Y] == segmentMatrix[n.X][n.Y] {
						queue = append(queue, neighbours[i])
						matrix[neighbours[i].X][neighbours[i].Y] = edgeDirection(Edge{neighbours[i], n, 0})
						opened[neighbours[i]] = true
					}
				}

			}
		}
	}
	return matrix
}


func DirectionMatrixToSegmentMatrixAndSegmentMap(directionMatrix [][]Direction) ([][]int, map[int][]Vertex) {
	// Initialize segment matrix
	segmentMatrix := make([][]int, pictureWidth)
	for x := range segmentMatrix {
		segmentMatrix[x] = make([]int, pictureHeight)
	}
	// Initialize disjoint sets
	djSets := make(map[Vertex]*Element)
	for x := range directionMatrix {
		for y := range directionMatrix[x] {
			set := MakeSet(Vertex{x, y})
			djSets[Vertex{x, y}] = set
		}
	}
	// Union nodes in same segment
	for x := range directionMatrix {
		for y := range directionMatrix[x] {
			node := Vertex{x, y}
			neighbour := nodeAndDirectionToNode(node, directionMatrix[x][y])
			s1, s2 := FindSet(djSets[node]), FindSet(djSets[neighbour])
			Union(s1, s2)
		}
	}
	// Create segments from sets
	setToSegment := make(map[*Element][]Vertex)
	for x := range directionMatrix {
		for y := range directionMatrix[x] {
			node := Vertex{x, y}
			set := FindSet(djSets[node])
			setToSegment[set] = append(setToSegment[set], node)
		}
	}
	segmentID := 0
	segmentMap := make(map[int][]Vertex)
	for set := range setToSegment {
		for i := range setToSegment[set] {
			node := setToSegment[set][i]
			segmentMatrix[node.X][node.Y] = segmentID
		}
		segmentMap[segmentID] = setToSegment[set]
		segmentID++
	}
	return segmentMatrix, segmentMap
}

func nodeAndDirectionToNode(node Vertex, direction Direction) Vertex {
	var n Vertex
	switch direction {
	case Up:
		n = Vertex{node.X, node.Y - 1}
	case Down:
		n = Vertex{node.X, node.Y + 1}
	case Right:
		n = Vertex{node.X + 1, node.Y}
	case Left:
		n = Vertex{node.X - 1, node.Y}
	default:
		return Vertex{node.X, node.Y}
	}
	if n.isInRange() {
		return n
	} else {
		return node
	}
}

func (node *Vertex) isInRange() bool {
	return node.X >= 0 && node.X < pictureWidth && node.Y >= 0 && node.Y < pictureHeight
}