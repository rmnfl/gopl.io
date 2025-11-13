package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	bigImg := createImage(width*2, height*2)
	img := superSample(bigImg)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func createImage(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
}

func superSample(bigImg image.Image) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			px1 := bigImg.At(px*2, py*2)
			px2 := bigImg.At(px*2+1, py*2)
			px3 := bigImg.At(px*2, py*2+1)
			px4 := bigImg.At(px*2+1, py*2+1)

			newColor := avgColor(px1, px2, px3, px4)
			img.Set(px, py, newColor)
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func avgColor(colors ...color.Color) color.Color {
	var sumR, sumG, sumB, sumA uint32
	for _, c := range colors {
		r, g, b, a := c.RGBA()

		sumR += r >> 8
		sumG += g >> 8
		sumB += b >> 8
		sumA += a >> 8
	}

	avgR, avgG, avgB, avgA := sumR/4, sumG/4, sumB/4, sumA/4
	return color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), uint8(avgA)}
}
