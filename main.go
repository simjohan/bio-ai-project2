package main

import (
	"log"
	"time"
	//"./tester"
	"math/rand"

	"strconv"
)

// global vars
var pictureWidth, pictureHeight, elites, edgeWeight, deviationWeight, generations int
var pic *Picture
var mutationRate, crossoverRate float64

func main() {

	/* params */
	mutationRate = 0.5
	crossoverRate = 0.9
	elites = 1
	edgeWeight = 1
	deviationWeight = -1
	generations = 10

	//picture, _ := readImageFromFile("test5x3.jpg")
	picture, _ := readImageFromFile("353013/Test image.jpg")
	pic = &picture
	pictureWidth = picture.width
	pictureHeight = picture.height
	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	log.Println("Started")


	pop := Population{}
	pop.InitPopulation(5)
	WriteImage("images/output/before"+strconv.Itoa(len(pop.Individuals[0].SegmentMap))+".png", SegmentMatrixToImage(pop.Individuals[0].SegmentMatrix, false))

	pop.WeightedSum()
	//pop.nsga2()

	//tester.CalculatePRI("/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/tester/run.py", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/353013/", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/output/paretofront/")
	log.Println("Elapsed", time.Since(start))

}
