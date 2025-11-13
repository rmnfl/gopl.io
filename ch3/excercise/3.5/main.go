package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			smooth := float64(n) - math.Log(math.Log(cmplx.Abs(v)))/math.Log(2)
			t := smooth / float64(iterations)
			return colorizeYCbCr(t)
			// return colorizeRGB(t)
		}
	}
	return color.Black
}

func colorizeRGB(t float64) color.Color {
	// Map t ∈ [0,1] smoothly across a color gradient
	r := uint8(255 * (0.5 + 0.5*math.Sin(20*math.Pi*t)))
	g := uint8(255 * (0.5 + 0.5*math.Sin(20*math.Pi*t+2*math.Pi/3)))
	b := uint8(255 * (0.5 + 0.5*math.Sin(20*math.Pi*t+4*math.Pi/3)))
	return color.RGBA{r, g, b, 255}
}

func colorizeYCbCr(t float64) color.Color {
	// Brightness (Y) goes from dark to bright
	Y := uint8(255 * math.Pow(t, 0.2)) // boosts highlights

	// Chrominance channels (Cb, Cr) oscillate to create color variations
	Cb := uint8(128 + 127*math.Sin(6.2831*t))
	Cr := uint8(128 + 127*math.Cos(6.2831*t))

	return color.YCbCr{Y, Cb, Cr}
}
