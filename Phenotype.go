package main

type Phenotype struct {
	segments     	 [][]Vertex
	segmentIdMap 	 map[Vertex]int
	picture 	 	 *Picture
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

func (p Phenotype) OverallDeviation() float64 {
	deviation := 0.0
	for _, segment := range p.segments {
		centeroid := averageSegmentColor(segment, p.picture)
		p2 := rgbaToPixel(centeroid)
		for _, node := range segment {
			p1 := p.picture.pixels[node.X][node.Y]
			deviation += euclideanDistance(&p1, &p2)
		}
	}
	return deviation
}

func (p Phenotype) EdgeValue() float64 {
	edgeValue := 0.0

	for s := range p.segments {
		for _, vertex := range p.segments[s] {
			node := Vertex{vertex.X, vertex.Y}
			p1 := p.picture.pixels[node.X][node.Y]
			neighbours := getAllCardinalNeighbours(node)
			for _, neighbour := range neighbours {
				if inBounds(neighbour, p.picture) {
					if p.segmentIdMap[vertex] != p.segmentIdMap[neighbour] {
						p2 := p.picture.pixels[neighbour.X][neighbour.Y]
						edgeValue += euclideanDistance(&p1, &p2)
					}
				}
			}
		}
	}

	return edgeValue
}