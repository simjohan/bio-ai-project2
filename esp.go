package main
//
//import (
//	"fmt"
//	"math"
//	"sort"
//	"math/rand"
//)
//
//var _ = fmt.Println
//
//type Direction int
//
//const (
//	Up Direction = iota
//	Down
//	Right
//	Left
//	Self
//)
//
//
//type ImageGraph struct {
//	Nodes []Node
//	Edges []Edge
//}
//
//func (iGraph *ImageGraph) Init() {
//	nodes := make([]Node, imageWidth*imageHeight)
//	edges := make([]Edge, (imageWidth-1)*imageHeight+(imageHeight-1)*imageWidth)
//	nodeI := 0
//	edgeI := 0
//
//	for x := 0; x < imageWidth; x++ {
//		for y := 0; y < imageHeight; y++ {
//			node := Node{x, y}
//			nodes[nodeI] = node
//			nodeI += 1
//			right, rightErr, down, downErr := node.neighbours()
//			if rightErr == nil {
//				edges[edgeI] = Edge{node, right, costOfEdge(node, right)}
//				edgeI += 1
//
//			}
//			if downErr == nil {
//				edges[edgeI] = Edge{node, down, costOfEdge(node, down)}
//				edgeI += 1
//			}
//
//		}
//	}
//	sort.Sort(ByCost(edges))
//	iGraph.Nodes = nodes
//	iGraph.Edges = edges
//}
//
//func (iGraph *ImageGraph) MakeSegments(k int) [][]Node {
//	nodes := iGraph.Nodes
//	edges := iGraph.Edges
//	tree := make([]Edge, 0)
//	disjointSets := make(map[Node]*DisjointSet)
//	maxInternal := make(map[*DisjointSet]float64)
//	size := make(map[*DisjointSet]int)
//	for i := 0; i < len(nodes); i++ {
//		set := makeSet()
//		size[&set] = 1
//		disjointSets[nodes[i]] = &set
//	}
//
//	for _, edge := range edges {
//		fromSet, toSet := find(disjointSets[edge.From]), find(disjointSets[edge.To])
//		if fromSet != toSet {
//			mInt := math.Min(maxInternal[fromSet]+float64(k/size[fromSet]), maxInternal[toSet]+float64(k/size[toSet]))
//			if edge.Cost <= mInt {
//				tree = append(tree, edge)
//				size1, size2 := size[fromSet], size[toSet]
//				Union(disjointSets[edge.From], disjointSets[edge.To])
//				maxInternal[fromSet] = edge.Cost
//				size[fromSet] = size1 + size2
//
//			}
//		}
//	}
//
//	directions := make(map[Node]Direction)
//	parent := make(map[Node]Node)
//	for _, n := range iGraph.Nodes {
//		parent[n] = n
//	}
//	// Create adjacency list for entire graph
//	adjacencyList := make(map[Node][]Node)
//	for _, edge := range tree {
//		adjacencyList[edge.From] = append(adjacencyList[edge.From], edge.To)
//		adjacencyList[edge.To] = append(adjacencyList[edge.To], edge.From)
//	}
//
//	// Build all segments
//	setToSegment := make(map[*DisjointSet][]Node)
//	for _, k := range iGraph.Nodes {
//		setToSegment[find(disjointSets[k])] = append(setToSegment[find(disjointSets[k])], k)
//	}
//	segments := make([][]Node, 0)
//	for _, value := range setToSegment {
//		segments = append(segments, value)
//	}
//
//	for _, segment := range segments {
//		visited := make(map[Node]bool)
//		queue := []Node{segment[0]}
//		for len(queue) > 0 {
//			n := queue[0]
//			visited[n] = true
//			queue = queue[1:]
//			for _, neighbour := range adjacencyList[n] {
//				if _, visited := visited[neighbour]; !visited {
//					queue = append(queue, neighbour)
//					parent[neighbour] = n
//				}
//			}
//		}
//	}
//	for node := range parent {
//		directions[node] = directionOfEdge(Edge{node, parent[node], 0})
//	}
//
//	return genotypeToPhenotype(directions)
//
//}
//
//
//
//func genotypeToPhenotype(directions map[Node]Direction) [][]Node {
//	djSets := make(map[Node]*DisjointSet)
//	for k := range directions {
//		set := makeSet()
//		djSets[k] = &set
//	}
//
//	for node, direction := range directions {
//		neighbour := nodeAndDirectionToNode(node, direction)
//		s1, s2 := find(djSets[node]), find(djSets[neighbour])
//		Union(s1, s2)
//	}
//
//	// Create segments from MSF
//	setToSegment := make(map[*DisjointSet][]Node)
//	for k := range directions {
//		setToSegment[find(djSets[k])] = append(setToSegment[find(djSets[k])], k)
//	}
//	segments := make([][]Node, 0)
//	for _, value := range setToSegment {
//		segments = append(segments, value)
//	}
//	return segments
//}
//
//func directionOfEdge(edge Edge) Direction {
//	dX, dY := edge.To.X-edge.From.X, edge.To.Y-edge.From.Y
//	switch {
//	case dX >= 1:
//		return Right
//	case dX <= -1:
//		return Left
//	case dY >= 1:
//		return Down
//	case dY <= -1:
//		return Up
//	default:
//		return Self
//	}
//}
//
//func nodeAndDirectionToNode(node Node, direction Direction) Node {
//	var n Node
//	switch direction {
//	case Up:
//		n = Node{node.X, node.Y - 1}
//	case Down:
//		n = Node{node.X, node.Y + 1}
//	case Right:
//		n = Node{node.X + 1, node.Y}
//	case Left:
//		n = Node{node.X - 1, node.Y}
//	default:
//		return Node{node.X, node.Y}
//	}
//	if n.isInRange() {
//		return n
//	} else {
//		fmt.Println(node, "to", n, "NOT IN RANGE")
//		return node
//	}
//}
//
//func randomDirection() Direction {
//	switch rand.Intn(5) {
//	case 0:
//		return Up
//	case 1:
//		return Down
//	case 2:
//		return Right
//	case 3:
//		return Left
//	}
//	return Self
//}
//
//func costOfEdge(from, to Node) float64 {
//	return rgbDistance(&imagePixels[from.X][from.Y], &imagePixels[to.X][to.Y])
//}
