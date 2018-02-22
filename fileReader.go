package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func readImageFromFile(fileName string) ([][]Pixel, error) {
	imageFile, err := os.Open("images/3/" + fileName)

	if err != nil {
		fmt.Println("image not found...")
		return nil, err
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

	return pixels, nil

}

func main() {
	pixels, _ := readImageFromFile("Test Image.jpg")
	fmt.Println(pixels[0][0])
	fmt.Println(pixels[199][32])
	fmt.Println(pixels[22][32])
	fmt.Println(euclideanDistance(&pixels[1][2], &pixels[199][32]))
	//generateRandomGenotype(pixels)
	fmt.Println(makeWeightTable(pixels))
}
