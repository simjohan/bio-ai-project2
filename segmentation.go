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


func (g Graph) GraphSegmentation(k int) [][]Vertex {
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
	directions := make(map[Vertex]Direction)
	parent     := make(map[Vertex]Vertex)

	/* init all node parents to itself */
	for _, node := range g.Vertices {
		parent[node] = node
	}

	//neighbours := make(map[Vertex][]Vertex)
	//for _, edge := range mst {
	//	neighbours[edge.U] = append(neighbours[edge.U], neighbours[edge.V])
	//	neighbours[edge.V] = append(neighbours[edge.V], neighbours[edge.U])
	//
	//}
	neighbours := make(map[Vertex][]Vertex)
	for _, edge := range mst {
		//from := Point{edge.U.(Point).X, edge.U.(Point).Y}
		//to := Point{edge.V.(Point).X, edge.V.(Point).Y}
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
	}
	//fmt.Println(directions)
	//genoToPheno(directions)
	//return segments
	return genoToPheno(directions)
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

//func findSegmentBorder(segment []Vertex) []Vertex {
//
//	borders := make([]*map[int]int, 4)
//	colMin := make(map[int]int)
//	colMax := make(map[int]int)
//	rowMin := make(map[int]int)
//	rowMax := make(map[int]int)
//
//	borders[0] = &colMin
//	borders[1] = &colMax
//	borders[2] = &rowMin
//	borders[3] = &rowMax
//
//	for v := range segment {
//
//		col := segment[v].(Point).X
//		row := segment[v].(Point).Y
//
//		if colMin == nil || colMin[col] > row {
//			colMin[col] = row
//		}
//		if colMax == nil || colMax[col] < row {
//			colMax[col] = row
//		}
//		if rowMin == nil || rowMin[row] > col {
//			rowMin[row] = col
//		}
//		if rowMax == nil || rowMax[row] < col {
//			rowMax[row] = col
//		}
//
//	}
//
//	borderNodes := make([]Vertex, 0)
//
//	for e := range borders {
//		for k, v := range *borders[e] {
//			if e == 0 || e == 1 {
//				borderNodes = append(borderNodes, Point{k, v})
//			} else {
//				borderNodes = append(borderNodes, Point{v, k})
//			}
//		}
//	}
//
//	return borderNodes
//
//}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

func genoToPheno(directions map[Vertex]Direction) [][]Vertex {

	disjoint := make(map[Vertex]*Element)
	for v := range directions {
		simpleSet := MakeSet(disjoint[v])
		disjoint[v] = simpleSet
	}

	for v, dir := range directions {
		adjacent := getNeighbourFromDirection(v, dir)
		set1     := FindSet(disjoint[v])
		set2     := FindSet(disjoint[adjacent])
		Union(set1, set2)
	}

	toSegment := make(map[*Element][]Vertex)
	for s := range directions {
		toSegment[FindSet(disjoint[s])] = append(toSegment[FindSet(disjoint[s])], s)
	}

	segments := make([][]Vertex, 0)
	for _, v := range toSegment {
		segments = append(segments, v)
	}

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

func getAllCardinalNeighbours(from Vertex) {
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