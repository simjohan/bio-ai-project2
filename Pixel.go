package main

type Pixel struct {
	R, G, B, A uint8
}

func rgbaToPixel(r, g, b, a uint8) Pixel {
	return Pixel{ r, g, b, a}
}

