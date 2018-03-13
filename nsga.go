package main

import (
	"sort"
	"math"
)

//
//import (
//	"sort"
//	"math"
//)
//


func (population *Population) NonDominatedSort() {
	// Initialize font set
	fronts := make([][]*MatrixIndividual, 0)
	// Create the first front list
	fronts = append(fronts, make([]*MatrixIndividual, 0))

	for i := range population.Individuals {
		individual := &population.Individuals[i]
		dominatedByIndividual := make([]*MatrixIndividual, 0)
		numDominatingIndividual := 0
		for j := range population.Individuals {
			// Same individual
			if i == j {
				continue
			} else if individual.IsDominating(&population.Individuals[j]) {
				dominatedByIndividual = append(dominatedByIndividual, &population.Individuals[j])
			} else if population.Individuals[j].IsDominating(individual) {
				numDominatingIndividual += 1
			}
		}
		individual.Dominates = dominatedByIndividual
		individual.DominatedBy = numDominatingIndividual
		// If none dominates individual, add individual to first front set
		if numDominatingIndividual == 0 {
			individual.Rank = 1
			fronts[0] = append(fronts[0], individual)
		}
	}
	frontCounter := 0
	for len(fronts[frontCounter]) > 0 {
		nextFront := make([]*MatrixIndividual, 0)
		for i := range fronts[frontCounter] {
			individual := fronts[frontCounter][i]
			for j := range individual.Dominates {
				dominatedIndividual := individual.Dominates[j]
				dominatedIndividual.DominatedBy -= 1
				if dominatedIndividual.DominatedBy == 0 {
					dominatedIndividual.Rank = frontCounter + 2
					nextFront = append(nextFront, dominatedIndividual)
				}
			}
		}
		fronts = append(fronts, nextFront)
		frontCounter += 1
	}
	population.Fronts = fronts
}

func (population *Population) CrowdingDistance() {
	maxDev, minDev, maxEdge, minEdge := population.getMaxMinValues()
	deltaDev, deltaEdge := maxDev-minDev, maxEdge-minEdge
	for i := range population.Fronts {
		if len(population.Fronts[i]) == 0 {
			continue
		}
		for j := range population.Fronts[i] {
			population.Fronts[i][j].CrowdingDistance = 0
		}
		sort.Sort(byOverallDeviation(population.Fronts[i]))
		population.Fronts[i][0].CrowdingDistance = math.Inf(1)
		population.Fronts[i][len(population.Fronts[i])-1].CrowdingDistance = math.Inf(1)
		for k := 1; k < len(population.Fronts[i])-1; k++ {
			population.Fronts[i][k].CrowdingDistance = population.Fronts[i][k].CrowdingDistance + ((population.Fronts[i][k+1].overallDeviation - population.Fronts[i][k-1].overallDeviation) / deltaDev)
		}
		sort.Sort(byEdgeValue(population.Fronts[i]))
		population.Fronts[i][0].CrowdingDistance = math.Inf(1)
		population.Fronts[i][len(population.Fronts[i])-1].CrowdingDistance = math.Inf(1)
		for k := 1; k < len(population.Fronts[i])-1; k++ {
			population.Fronts[i][k].CrowdingDistance = population.Fronts[i][k].CrowdingDistance + ((population.Fronts[i][k-1].edgeValue - population.Fronts[i][k+1].edgeValue) / deltaEdge)
		}
	}
}
func (population *Population) getMaxMinValues() (float64, float64, float64, float64) {
	maxDev, maxEdge := 0.0, 0.0
	minDev, minEdge := math.Inf(1), math.Inf(1)
	for i := range population.Individuals {
		if population.Individuals[i].edgeValue> maxEdge {
			maxEdge = population.Individuals[i].edgeValue
		}
		if population.Individuals[i].edgeValue < minEdge {
			minEdge = population.Individuals[i].edgeValue
		}
		if population.Individuals[i].overallDeviation > maxDev {
			maxDev = population.Individuals[i].overallDeviation
		}
		if population.Individuals[i].overallDeviation < minDev {
			minDev = population.Individuals[i].overallDeviation
		}
	}
	return maxDev, minDev, maxEdge, minEdge
}


//
//
//func (p *Population) NonDominatedSort() {
//	/* Map to store all individuals dominated by the key Individual*/
//	dominatedBy := make(map[Individual][]Individual)
//
//	/* Map to store how many individuals dominate the key */
//	dominationCount := make(map[Individual]int)
//
//	front := make(map[int][]Individual)
//	rank  := make(map[Individual]int)
//
//	/* initialize */
//	for _, i1 := range population.Individuals {
//		dominatedBy[i1] = nil
//		dominationCount[i1] = 0
//
//		for _, i2 := range population.Individuals {
//			/* if i1==i2, skip */
//			//if i1.Fitness == i2.Fitness { // TODO: this will not work, need to find another way do chech for similarity
//			//	continue
//			//}
//			// if i1 dominates i2
//			if dominates(&i1, &i2) {
//				dominatedBy[i1] = append(dominatedBy[i1], i2)
//			} else if dominates(&i2, &i1) {
//				dominationCount[i1]++
//			}
//		}
//
//		if dominationCount[i1] == 0 {
//			front[0] = append(front[0], i1)
//			rank[i1] = 1
//		}
//	}
//
//	i := 1
//	for len(front[i]) > 0 {
//		queue := make([]Individual, 0)
//		for _, i1 := range front[i] {
//			for _, i2 := range dominatedBy[i1] {
//				dominationCount[i2]--
//				if dominationCount[i2] == 0 {
//					rank[i2] = i + 1
//					queue = append(queue, i2)
//				}
//			}
//		}
//		i++
//		front[i] = append(front[i], queue...)
//	}
//
//	return front
//}
//
//func crowdingDistance(front []Individual) {
//
//	/* calculate for OverallDeviation */
//	sort.Sort(byOverallDeviation(front))
//	//for _, individual := range fron
//	front[0].crowdingDistance, front[len(front)-1].crowdingDistance = math.Inf(0), math.Inf(0)
//	for i := 1; i < len(front); i++ {
//		front[i].crowdingDistance =
//			front[i].crowdingDistance +
//				(front[i + 1].overallDeviation - front[i - 1].overallDeviation) /
//					(front[len(front)-1].overallDeviation - front[0].overallDeviation)
//	}
//
//
//	/* calculate for EdgeValue */
//	sort.Sort(byEdgeValue(front))
//	front[0].crowdingDistance, front[len(front)-1].crowdingDistance = math.Inf(0), math.Inf(0)
//	for i := 1; i < len(front); i++ {
//		front[i].crowdingDistance =
//			front[i].crowdingDistance +
//				(front[i + 1].edgeValue - front[i - 1].edgeValue) /
//					(front[0].edgeValue - front[len(front)-1].edgeValue)
//	}
//}
//
//
/* Returns true if i1 dominates i2 */ // TODO: overallDeviation and edgeValue is currently not set anywhere
func dominates(i1, i2 *MatrixIndividual) bool {
	if i1.overallDeviation < i2.overallDeviation &&
		i1.edgeValue > i2.edgeValue {
		return true
	}
	return false
}

/* sorting implementations for objective functions */
/* overallDeviation is minimized */
type byOverallDeviation 						[]*MatrixIndividual
func (f byOverallDeviation) Len() int 			{ return len(f) }
func (f byOverallDeviation) Swap(i, j int) 		{ f[i], f[j] = f[j], f[i] }
func (f byOverallDeviation) Less(i, j int) bool { return f[i].overallDeviation < f[j].overallDeviation }

/* edge is maximized */
type byEdgeValue 								[]*MatrixIndividual
func (f byEdgeValue) Len() int 					{ return len(f) }
func (f byEdgeValue) Swap(i, j int) 			{ f[i], f[j] = f[j], f[i] }
func (f byEdgeValue) Less(i, j int) bool 		{ return f[i].edgeValue > f[j].edgeValue }

/* fitness is maximized */
type byFitness	 								[]MatrixIndividual
func (f byFitness) Len() int 					{ return len(f) }
func (f byFitness) Swap(i, j int) 				{ f[i], f[j] = f[j], f[i] }
func (f byFitness) Less(i, j int) bool 			{ return f[i].Fitness > f[j].Fitness }



///* partial order */
//type byPartialOrder
type ByNSGA []MatrixIndividual

func (p ByNSGA) Len() int      { return len(p) }
func (p ByNSGA) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByNSGA) Less(i, j int) bool {
	if p[i].Rank < p[j].Rank {
		return true
	} else if p[i].Rank == p[j].Rank && p[i].CrowdingDistance > p[j].CrowdingDistance {
		return true
	} else {
		return false
	}
}