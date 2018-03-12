package main

import (
	"math/rand"
)

type Genotype struct {
	genes []Direction
}

func (g *Genotype) mutate() {
	// Selecet a random number between 50-500 to undergo mutation
	genesToMutate := rand.Intn((2000-50) + 50)
	selectedGenes := rand.Perm(len(g.genes))[0:genesToMutate]

	for gene := range selectedGenes {
		g.genes[gene] = generateRandomDirection()
	}

}


func generateRandomDirection() Direction {
	switch rand.Intn(5) {
	case 0:
		return Up
	case 1:
		return Down
	case 2:
		return Left
	case 3:
		return Right
	case 4:
		return None
	}
	return None
}