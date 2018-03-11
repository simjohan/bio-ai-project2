package main

import "sort"

func nonDominatedSort(population *Population) map[int][]Individual {
	/* Map to store all individuals dominated by the key Individual*/
	dominatedBy := make(map[Individual][]Individual)

	/* Map to store how many individuals dominate the key */
	dominationCount := make(map[Individual]int)

	front := make(map[int][]Individual)
	rank  := make(map[Individual]int)

	/* initialize */
	for _, i1 := range population.Individuals {
		dominatedBy[i1] = nil
		dominationCount[i1] = 0

		for _, i2 := range population.Individuals {
			/* if i1==i2, skip */
			if i1.Fitness == i2.Fitness { // TODO: this will not work, need to find another way do chech for similarity
				continue
			}
			if dominates(&i1, &i2) {
				dominatedBy[i1] = append(dominatedBy[i1], i2)
			} else if dominates(&i2, &i1) {
				dominationCount[i1]++
			}
		}

		if dominationCount[i1] == 0 {
			front[0] = append(front[0], i1)
			rank[i1] = 1
		}
	}

	i := 1
	for len(front[i]) > 0 {
		queue := make([]Individual, 0)
		for _, i1 := range front[i] {
			for _, i2 := range dominatedBy[i1] {
				dominationCount[i2]--
				if dominationCount[i2] == 0 {
					rank[i2] = i + 1
					queue = append(queue, i2)
				}
			}
		}
		i++
		front[i] = append(front[i], queue...)
	}

	return front
}

func crowdingDistance(front []Individual) {

	distances := make(map[Individual]float64)

	for _, individual := range front {
		distances[individual] = 0.0
	}

	/* calculate for OverallDeviation */
	sort.Sort(byOverallDeviation(front))


	/* calculate for EdgeValue */
	sort.Sort(byEdgeValue(front))

}


/* Returns true if i1 dominates i2 */ // TODO: overallDeviation and edgeValue is currently not set anywhere
func dominates(i1, i2 *Individual) bool {
	if i1.overallDeviation < i2.overallDeviation &&
		i1.edgeValue > i2.edgeValue {
		return true
	}
	return false
}

/* sorting implementations for objective functions */
/* edge is minimized */
type byOverallDeviation 						[]Individual
func (f byOverallDeviation) Len() int 			{ return len(f) }
func (f byOverallDeviation) Swap(i, j int) 		{ f[i], f[j] = f[j], f[i] }
func (f byOverallDeviation) Less(i, j int) bool { return f[i].overallDeviation < f[j].overallDeviation }

/* edge is maximized */
type byEdgeValue 								[]Individual
func (f byEdgeValue) Len() int 					{ return len(f) }
func (f byEdgeValue) Swap(i, j int) 			{ f[i], f[j] = f[j], f[i] }
func (f byEdgeValue) Less(i, j int) bool 		{ return f[i].edgeValue > f[j].edgeValue }