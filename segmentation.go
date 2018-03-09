package main

import (
	"sort"
	"math"
	"fmt"
)

type Edge struct {
	U, V   Vertex
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f\n", e.U, e.V, e.Weight)
}

type Vertex = interface{}

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}

type Point struct {
	X, Y int
}


func (g Graph) GraphSegmentation(k int) [][]Vertex {
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*Element)
	maxInternal := make(map[*Element]float64)
	sizeMap := make(map[*Element]int)

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
				size1, size2 := sizeMap[from], sizeMap[to]
				Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
			}
		}
	}

	segMap := make(map[*Element][]Vertex)
	for _,vertex := range g.Vertices {
		rep := FindSet(disjointMap[vertex])
		segMap[rep] = append(segMap[rep], vertex)
	}

	segments := make([][]Vertex,0)
	for _, value := range segMap {
		segments = append(segments, value)
	}
	return segments
}