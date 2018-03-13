package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"image/color"
	"image/png"
	"log"
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

func inBounds(node Vertex) bool {
	if node.X >= 0 && node.X < pictureWidth && node.Y >=0 && node.Y < pictureHeight {
		return true
	}
	return false
}

// Outputs a JPEG image of the image to the path
func WriteImage(path string, image image.Image) {
	log.Println("Writing", "\""+path+"\"")

	outfile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	pngenc := png.Encoder{CompressionLevel: -1}
	//jpeg.Encode(outfile, image, &jpeg.Options{Quality:100})
	pngenc.Encode(outfile, image)
}

func SegmentMatrixToImage(matrix [][]int, bnw bool) image.Image {
	pixels := make([][]Pixel, pictureWidth)
	for i := range pixels {
		pixels[i] = make([]Pixel, pictureHeight)
	}
	for x := range matrix {
		for y := range matrix[x] {
			n := Vertex{x, y}
			right, errRight, down, downErr := getNeighbours(n)
			border := false
			if errRight == nil {
				if matrix[x][y] != matrix[right.X][right.Y] {
					border = true
				}
			}
			if downErr == nil {
				if matrix[x][y] != matrix[down.X][down.Y] {
					border = true
				}
			}
			if border {
				if bnw {
					pixels[n.X][n.Y] = Pixel{0, 0, 0, 255}
				} else {
					pixels[n.X][n.Y] = Pixel{0, 255, 0, 255}
				}
			} else {
				if bnw {
					pixels[n.X][n.Y] = Pixel{255, 255, 255, 255}
				} else {
					pixels[n.X][n.Y] = pic.pixels[n.X][n.Y]
				}
			}

		}
	}
	return PixelArrayToRgbaImage(pixels)
}

func PixelArrayToRgbaImage(pixels [][]Pixel) image.Image {
	width, height := len(pixels), len(pixels[0])
	rgbaImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := pixels[x][y]
			rgbaImage.Set(x, y, color.RGBA{p.R, p.G, p.B, p.A})
		}
	}
	return rgbaImage
}


func drawGroundTruthPicture(picture *Picture, segments [][]Vertex, segMatrix [][]int, name string) {
	var img = image.NewRGBA(image.Rect(0,0, picture.width, picture.height))
	for s := range segments {
		for _, vertex := range segments[s]{
			white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
			img.Set(vertex.X, vertex.Y, white)

			node := Vertex{vertex.X, vertex.Y}
			neighbours := getTwoCardinalNeighbours(node)
			for _, neighbour := range neighbours {
				if inBounds(neighbour) {
					if segMatrix[node.X][node.Y] != segMatrix[neighbour.X][neighbour.Y] {
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