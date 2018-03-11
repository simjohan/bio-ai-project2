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
			pixels[x][y] = rgbaToPixel(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
		}
	}

	picture = Picture{pixels, width, height, width*height}
	return picture, nil

}

func drawPicture(picture Picture) {
	var img = image.NewRGBA(image.Rect(0,0, picture.width, picture.height))
	for x := 0; x < picture.width; x++ {
		for y := 0; y < picture.height; y++ {
			pixel := picture.pixels[x][y]
			rgb := color.RGBA{R: pixel.R, G: pixel.G, B: pixel.B, A: pixel.A}
			//fmt.Println(rgb)
			img.Set(x, y, rgb)
		}
	}
	f, err := os.Create("images/output/draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
