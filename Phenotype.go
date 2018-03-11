package main

type Phenotype struct {
	segments [][]Vertex
	segmentIdMap map[Vertex]int
}

func (p Phenotype) Init(picture *Picture, segs [][]Vertex) Phenotype {

	return Phenotype{segs, generateSegmentIdMap(segs)}
}

func generateSegmentIdMap(segments [][]Vertex) map[Vertex]int {
	segmentIdMap := make(map[Vertex]int)
	for s := range segments {
		for v := range segments[s] {
			segmentIdMap[segments[s][v]] = s
		}
	}
	return segmentIdMap
}

