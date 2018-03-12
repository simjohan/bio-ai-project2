package main

import (
	"math/rand"

	"sort"
	"fmt"
)

type Individual struct {
	Genotype  		 Genotype
	Phenotype 		 Phenotype
	Fitness   		 float64
	overallDeviation float64
	edgeValue 		 float64
	crowdingDistance float64
	Rank			 int
}

type Population struct {
	Individuals []Individual
}

//func MOEA(picture *Picture) {
//
//	generations := 10
//	populationSize := 10
//
//	initialPopulation := initPopulation(picture, populationSize)
//	fronts := nonDominatedSort(&initialPopulation)
//
//	for g := 0; g < generations; g++ {
//
//		nextPopulation := Population{[]Individual{} }
//		for len(nextPopulation.Individuals) <= populationSize {
//
//		}
//	}
//
//
//}



func (p *Population) uniformCrossover(parent1, parent2 *Individual, picture *Picture) (Individual, Individual) {

	var offspring1, offspring2 Individual

	numGenes := len(parent1.Genotype.genes)
	crossoverPoints := rand.Perm(numGenes)[0:int(numGenes/2)]
	sort.Ints(crossoverPoints)

	genes1, genes2 := make([]Direction, 0), make([]Direction, 0)

	prev := 0
	parents := []Individual{*parent1, *parent2}
	for i, n := range crossoverPoints {
		if i == len(crossoverPoints) - 1 {
			n = numGenes
		}
		genes1 = append(genes1, parents[0].Genotype.genes[prev:n]...)
		genes2 = append(genes2, parents[1].Genotype.genes[prev:n]...)
		parents[0], parents[1] = parents[1], parents[0]
		prev = n
	}

	fmt.Println(genes1)
	fmt.Println(genes2)

	pheno1 := initPhenotype(genoToPheno(genes1), picture)
	pheno2 := initPhenotype(genoToPheno(genes2), picture)
	offspring1 = Individual{Genotype{genes1}, pheno1, 0, pheno1.OverallDeviation(), pheno1.EdgeValue(), 0.0, 0}
	offspring2 = Individual{Genotype{genes2}, pheno2, 0, pheno2.OverallDeviation(), pheno2.EdgeValue(), 0.0, 0}

	return offspring1, offspring2

}

// this is shit
func (i *Individual) calculateWightedFitness() float64 {
	return i.Phenotype.EdgeValue() + i.Phenotype.OverallDeviation()*-1
}

//func initPopulation(picture *Picture, populationSize int) Population {
//
//	log.Println("Initializing population (size", strconv.Itoa(populationSize) + ")")
//
//	imageGraph := makeGraph(picture)
//
//	individuals := make([]Individual, 0)
//	for i := 0; i < populationSize; i++ {
//		// random value from 200-6000
//		randomKValue := rand.Intn((6000 - 200) + 200)
//
//		pheno, geno := imageGraph.GraphSegmentation(randomKValue)
//		genotype := Genotype{geno}
//		phenotype := initPhenotype(pheno, picture)
//
//		individual := Individual{genotype, phenotype, 0, phenotype.OverallDeviation(), phenotype.EdgeValue(), 0.0, 0}
//		individuals = append(individuals, individual)
//
//		log.Println("Individual", i+1, "/", populationSize)
//	}
//
//	return Population{individuals}
//
//}

