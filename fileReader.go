package main

import (
	"fmt"
	"image/jpeg"
	"image"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func readImageFromFile(fileName string) ([][]Pixel, error) {
	imageFile, err := os.Open("images/1/" + fileName)

	if err != nil {
		fmt.Println("image not found...")
		return nil, err
	}

	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel

	for y := 0; y < height; y++{
		var row []Pixel
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(0, 0).RGBA()
			row = append(row, rgbaToPixel(uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil

}

func main() {
	pixels, _ := readImageFromFile("Test Image.jpg")
	fmt.Println(pixels[0][2])


}
