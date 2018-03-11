package main

import "math/rand"

type Individual struct {
	genotype  Genotype
	phenotype Phenotype
	fitness   float64
}

type Population struct {
	Individuals []Individual
}

func initPopulation(picture *Picture, populationSize int) Population {

	individuals := make([]Individual, 0)
	for i := 0; i < populationSize; i++ {
		// random value from 200-2000
		randomKValue := rand.Intn((2000 - 200) + 200)

		individual :=

	}

}