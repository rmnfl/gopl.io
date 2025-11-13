package bigfloat

import (
	"image"
	"image/color"
	"math/big"
)

func Render(xmin, xmax, ymin, ymax float64, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin

			xFloat := new(big.Float).SetFloat64(x)
			yFloat := new(big.Float).SetFloat64(y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(xFloat, yFloat))
		}
	}
	return img
}

func mandelbrot(zRe *big.Float, zIm *big.Float) color.Color {
	const iterations = 100
	const contrast = 15

	var vRe, vIm = new(big.Float), new(big.Float)

	for n := uint8(0); n < iterations; n++ {
		// Compute zx^2, zy^2, and |z|^2
		vRe2 := new(big.Float).Mul(vRe, vRe)
		vIm2 := new(big.Float).Mul(vIm, vIm)
		abs2 := new(big.Float).Add(vRe2, vIm2)

		if abs2.Cmp(new(big.Float).SetInt64(4)) > 0 {
			return color.Gray{255 - contrast*n}
		}

		// Compute temporary 2*zx*zy
		vImNew := new(big.Float).Mul(vRe, vIm)
		vImNew.Mul(vImNew, new(big.Float).SetInt64(2))
		vImNew.Add(vImNew, zIm)

		// zx = zx^2 - zy^2 + x0
		vRe.Sub(vRe2, vIm2)
		vRe.Add(vRe, zRe)

		vIm.Set(vImNew)
	}
	return color.Black
}
