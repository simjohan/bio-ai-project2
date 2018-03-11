package main

import (
	"math/rand"
	"log"
	"strconv"
	"sort"
	"fmt"
)

type Individual struct {
	Genotype  Genotype
	Phenotype Phenotype
	Fitness   float64
	overallDeviation float64
	edgeValue 		 float64
}

type Population struct {
	Individuals []Individual
}

func (p Population) MOEA() {

	generations := 10
	var best Individual
	var worst Individual

	for g := 0; g < generations; g++ {

		p.Individuals[g].Fitness = p.Individuals[g].calculateWightedFitness()
		fmt.Println(p.Individuals[g].Fitness)

	}

	sort.Slice(p.Individuals, func(i, j int) bool {
		return p.Individuals[i].Fitness > p.Individuals[j].Fitness
	})

	best = p.Individuals[0]
	worst = p.Individuals[9]

	drawGroundTruthPicture(best.Phenotype.picture, best.Phenotype.segments, best.Phenotype.segmentIdMap, "best")
	drawGroundTruthPicture(worst.Phenotype.picture, worst.Phenotype.segments, worst.Phenotype.segmentIdMap, "worst")

}

func (i *Individual) calculateWightedFitness() float64 {
	return i.Phenotype.EdgeValue() + i.Phenotype.OverallDeviation()*-1
}

func initPopulation(picture *Picture, populationSize int) Population {

	log.Println("Initializing population (size", strconv.Itoa(populationSize) + ")")

	imageGraph := makeGraph(picture)

	individuals := make([]Individual, 0)
	for i := 0; i < populationSize; i++ {
		// random value from 200-6000
		randomKValue := rand.Intn((6000 - 200) + 200)

		pheno, geno := imageGraph.GraphSegmentation(randomKValue)
		genotype := Genotype{geno}
		phenotype := Phenotype{pheno, generateSegmentIdMap(pheno), picture}

		individual := Individual{genotype, phenotype, 0, phenotype.OverallDeviation(), phenotype.EdgeValue()}
		individuals = append(individuals, individual)

		log.Println("Individual", i+1, "/", populationSize)
	}

	return Population{individuals}

}

