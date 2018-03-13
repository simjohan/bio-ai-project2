package main

import (
	"log"
	"time"
	"./tester"
	"math/rand"

)

// global vars
var pictureWidth, pictureHeight, elites, edgeWeight, deviationWeight, generations int
var pic *Picture
var mutationRate, crossoverRate float64

func main() {

	/* params */
	mutationRate = 0.6
	crossoverRate = 1
	elites = 1
	edgeWeight = 1
	deviationWeight = -1
	generations = 6

	//picture, _ := readImageFromFile("test5x3.jpg")
	picture, _ := readImageFromFile("demo/test image_3.jpg")
	pic = &picture
	pictureWidth = picture.width
	pictureHeight = picture.height
	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	log.Println("Started")


	pop := Population{}
	pop.InitPopulation(15)
	////WriteImage("images/output/before"+strconv.Itoa(len(pop.Individuals[0].SegmentMap))+".png", SegmentMatrixToImage(pop.Individuals[0].SegmentMatrix, false))
	////
	////pop.WeightedSum()
	pop.nsga2()

	tester.CalculatePRI("/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/tester/run.py", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/demo/GT/3/", "/Users/simenjohansen/Documents/skole/Bio-AI/bio-ai-project2/images/output/paretofront/")
	log.Println("Elapsed", time.Since(start))

}

func getBelowAvgSegments(segMap map[int][]Vertex) []int{
	sum := 0
	for _, seg := range segMap {
		sum += len(seg)
	}
	avg := sum / len(segMap)
	res := make([]int, 0)
	for k, seg := range segMap {
		if len(seg) < avg {
			res = append(res, k)
		}
	}
	return res
}
