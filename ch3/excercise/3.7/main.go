package main

import (
	"image"
	"image/color"
	"image/png"
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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(z) = z^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	roots := []complex128{1 + 0i, 0 + 1i, -1 + 0i, 0 - 1i}
	colors := []color.RGBA{
		{255, 0, 0, 255},   // 1 → red
		{0, 255, 0, 255},   // i → green
		{0, 0, 255, 255},   // -1 → blue
		{255, 255, 0, 255}, // -i → yellow
	}

	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4

		for index, root := range roots {
			if cmplx.Abs(z-root) < 1e-6 {
				base := colors[index]
				shade := 255 - contrast*i
				return color.RGBA{
					R: uint8(float64(base.R) * float64(shade) / 255),
					G: uint8(float64(base.G) * float64(shade) / 255),
					B: uint8(float64(base.B) * float64(shade) / 255),
					A: 255,
				}
			}
		}
	}
	return color.Black
}
