package main

import (
	"log"
	"time"
	//"./tester"

	"math/rand"
	"fmt"
)

// global vars
var pictureWidth, pictureHeight int

func main() {
	picture, _ := readImageFromFile("test5x3.jpg")
	//picture, _ := readImageFromFile("147091/Test image.jpg")
	pictureWidth = picture.width
	pictureHeight = picture.height
	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	log.Println("Started")

	imageGraph := makeGraph(&picture)
	pheno, geno := imageGraph.GraphSegmentation(2)
	t := initNewIndividual(pheno, geno )
	fmt.Println(t.DirectionMatrix)


	//tester.CalculatePRI("/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/tester/run.py", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/147091/", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/output/")
	log.Println("Elapsed", time.Since(start))

}
//
//func testGenoToPheno() {
//	dirs := make(map[Vertex]Direction)
//	dirs[Vertex{0, 0}] = Right
//	dirs[Vertex{1, 0}] = Down
//	dirs[Vertex{2, 0}] = Left
//	dirs[Vertex{3, 0}] = None
//	dirs[Vertex{0, 1}] = Up
//	dirs[Vertex{1, 1}] = Left
//	dirs[Vertex{2, 1}] = Up
//	dirs[Vertex{3, 1}] = Up
//	dirs[Vertex{0, 2}] = Right
//	dirs[Vertex{1, 2}] = Right
//	dirs[Vertex{2, 2}] = Down
//	dirs[Vertex{3, 2}] = Up
//	dirs[Vertex{0, 3}] = Up
//	dirs[Vertex{1, 3}] = Left
//	dirs[Vertex{2, 3}] = Right
//	dirs[Vertex{3, 3}] = None
//
//	fmt.Println(genoToPheno(dirs))
//}
