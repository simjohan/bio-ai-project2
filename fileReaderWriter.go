package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"image/color"
	"image/png"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func readImageFromFile(fileName string) (Picture, error) {
	imageFile, err := os.Open("images/" + fileName)
	var picture Picture
	if err != nil {
		fmt.Println("image not found...")
		return picture, err
	}

	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	pixels := make([][]Pixel, width)
	//var pixels [][]Pixel
	for i := 0; i < width; i++ {
		pixels[i] = make([]Pixel, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[x][y] = colorToPixel(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}

	picture = Picture{pixels, width, height, width*height}
	return picture, nil

}

func drawPicture(picture *Picture, name string) {
	var img = image.NewRGBA(image.Rect(0,0, picture.width, picture.height))
	for x := 0; x < picture.width; x++ {
		for y := 0; y < picture.height; y++ {
			pixel := picture.pixels[x][y]
			rgb := color.RGBA{R: pixel.R, G: pixel.G, B: pixel.B, A: pixel.A}
			//fmt.Println(rgb)
			img.Set(x, y, rgb)
		}
	}
	f, err := os.Create("images/output/draw"+name+".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func inBounds(node Vertex, picture *Picture) bool {
	if node.X >= 0 && node.X < pictureWidth && node.Y >=0 && node.Y < pictureHeight {
		return true
	}
	return false
}

func drawGroundTruthPicture(picture *Picture, segments [][]Vertex, segmentIdMap map[Vertex]int, name string) {
	var img = image.NewRGBA(image.Rect(0,0, picture.width, picture.height))
	for s := range segments {
		for _, vertex := range segments[s]{
			white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
			img.Set(vertex.X, vertex.Y, white)

			node := Vertex{vertex.X, vertex.Y}
			neighbours := getAllCardinalNeighbours(node)
			for _, neighbour := range neighbours {
				if inBounds(neighbour, picture) {
					if segmentIdMap[vertex] != segmentIdMap[neighbour] {
						black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
						img.Set(vertex.X, vertex.Y, black)
					}
				}
			}

		}
	}
	f, err := os.Create("images/output/draw_"+name+".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}