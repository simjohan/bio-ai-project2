package main

import (
	"math/rand"
	"log"
	"strconv"
	"sort"
	"math"
)


type Population struct {
	Individuals []MatrixIndividual
	Size int
	GenerationNumber int
	Fronts [][]*MatrixIndividual
}

func (population *Population) nsga2() {

	for g := 0; g < generations; g++ {
		for i := 0; i < population.Size/2; i++ {
			tournamentIndexs := rand.Perm(population.Size)
			parent1, parent2 := tournamentIndexs[0:2], tournamentIndexs[0:2][2:4]
			parent1Index := int(math.Min(float64(parent1[0]), float64(parent1[1])))
			parent2Index := int(math.Min(float64(parent2[0]), float64(parent2[1])))
			child1, child2 := Crossover(population.Individuals[parent1Index], population.Individuals[parent2Index])
			population.Individuals = append(population.Individuals, child1, child2)
		}

		population.NonDominatedSort()
		population.CrowdingDistance()
		sort.Sort(ByNSGA(population.Individuals))
		population.Individuals = population.Individuals[:population.Size]
		population.NonDominatedSort()
		population.CrowdingDistance()
		sort.Sort(ByNSGA(population.Individuals))


		log.Println("Generation:", g, "edgeValue", population.Individuals[0].edgeValue, "deviation", population.Individuals[0].overallDeviation, "Number of fronts:", len(population.Fronts), "Size of front 1:", len(population.Fronts[0]))
	}

	for i := range population.Fronts[0] {

		WriteImage("images/output/paretofront/"+strconv.Itoa(len(population.Fronts[0][i].SegmentMap))+"_"+strconv.Itoa(i)+".png",
			SegmentMatrixToImage(population.Fronts[0][i].SegmentMatrix, true))
		WriteImage("images/output/paretofront_colored/"+strconv.Itoa(len(population.Fronts[0][i].SegmentMap))+"_"+strconv.Itoa(i)+".png",
			SegmentMatrixToImage(population.Fronts[0][i].SegmentMatrix, false))
	}

}

func (population *Population) WeightedSum() {
	WriteImage("images/output/before.png", SegmentMatrixToImage(population.Individuals[0].SegmentMatrix, false))
	for g := 0; g <= generations; g++ {

		sort.Sort(byFitness(population.Individuals))


		newPop := make([]MatrixIndividual, 0)

		// elitism
		newPop = append(newPop, population.Individuals[0:elites]...)

		for fh := 0; fh < population.Size / 2; fh++ {
			tournamentIndexs := rand.Perm(population.Size)[0:4]
			tournament := make([]MatrixIndividual, 0)

			for _, i := range tournamentIndexs {
				tournament = append(tournament, population.Individuals[i])
			}
			sort.Sort(byFitness(tournament))
			child1, child2 := Crossover(tournament[0], tournament[1])
			newPop = append(newPop, child1, child2)
		}

		population.Individuals = newPop[:population.Size]
		sort.Sort(byFitness(population.Individuals))

		log.Println("Generation", g, "edge:", population.Individuals[1].edgeValue, "Best:", population.Individuals[0].Fitness, "Worst:", population.Individuals[len(population.Individuals)-1].Fitness)
	}
	WriteImage("images/output/weighted_sum/"+strconv.Itoa(len(population.Individuals[0].SegmentMap))+".png",
		SegmentMatrixToImage(population.Individuals[0].SegmentMatrix, true))
	WriteImage("images/output/weighted_sum_colored/"+strconv.Itoa(len(population.Individuals[0].SegmentMap))+".png",
		SegmentMatrixToImage(population.Individuals[0].SegmentMatrix, false))
}


func (population *Population) InitPopulation(populationSize int)  {

	log.Println("Initializing population (size", strconv.Itoa(populationSize) + ")")

	imageGraph := makeGraph(pic)
	population.Size = populationSize

	for i := 0; i < populationSize; i++ {

		individual := MatrixIndividual{}
		individual.Init(imageGraph)
		individual.edgeValue = individual.EdgeValue()
		individual.overallDeviation = individual.OverallDeviation()

		population.Individuals = append(population.Individuals, individual)

		log.Println("Individual", i+1, "/", populationSize)
	}

}


func Crossover(parent1, parent2 MatrixIndividual) (MatrixIndividual, MatrixIndividual) {
	child1, child2 := MatrixIndividual{}, MatrixIndividual{}
	child1.DirectionMatrix = parent1.DirectionMatrix
	child2.DirectionMatrix = parent2.DirectionMatrix
	// Choose random segment from parent1 to add to child2
	if rand.Float64() < crossoverRate {
		for i := 0; i < 1; i++ {
			id := randomSegmentId(parent1.SegmentMap)
			for _, n := range parent1.SegmentMap[id] {
				child2.DirectionMatrix[n.X][n.Y] = parent1.DirectionMatrix[n.X][n.Y]
			}
			// Choose random segment from parent2 to add to child1
			id = randomSegmentId(parent2.SegmentMap)
			for _, n := range parent2.SegmentMap[id] {
				child1.DirectionMatrix[n.X][n.Y] = parent2.DirectionMatrix[n.X][n.Y]
			}

		}
	}
	child1.SegmentMatrix, child1.SegmentMap = DirectionMatrixToSegmentMatrixAndSegmentMap(child1.DirectionMatrix)
	child2.SegmentMatrix, child2.SegmentMap = DirectionMatrixToSegmentMatrixAndSegmentMap(child2.DirectionMatrix)
	if rand.Float64() < mutationRate {
		for i := 0; i < rand.Intn(100-2)+2; i++ {
			child1.mutate()
			child2.mutate()
		}
		child1.DirectionMatrix = SegmentMatrixAndSegmentMapToDirectionMatrix(child1.SegmentMatrix, child1.SegmentMap)
		child2.DirectionMatrix = SegmentMatrixAndSegmentMapToDirectionMatrix(child2.SegmentMatrix, child2.SegmentMap)
	}
	child1.CalculateFitness()
	child2.CalculateFitness()
	return child1, child2
}


