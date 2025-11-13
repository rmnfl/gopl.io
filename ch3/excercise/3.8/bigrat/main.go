// package bigrat

// import (
// 	"image"
// 	"image/color"
// 	"math/big"
// )

// var (
// 	four = new(big.Rat).SetInt64(4)
// 	two  = new(big.Rat).SetInt64(2)
// )

// func Render(xmin, xmax, ymin, ymax float64, width, height int) image.Image {
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))

// 	// Pre-calculate ratios to avoid repeated calculations
// 	xRatio := (xmax - xmin) / float64(width)
// 	yRatio := (ymax - ymin) / float64(height)

// 	// Reuse rat objects for pixel coordinates
// 	xRat := new(big.Rat)
// 	yRat := new(big.Rat)

// 	for py := 0; py < height; py++ {
// 		y := float64(py)*yRatio + ymin
// 		yRat.SetFloat64(y)

// 		for px := 0; px < width; px++ {
// 			x := float64(px)*xRatio + xmin
// 			xRat.SetFloat64(x)

// 			// Image point (px, py) represents complex value z.
// 			img.Set(px, py, mandelbrot(xRat, yRat))
// 		}
// 	}
// 	return img
// }

// func mandelbrot(zRe *big.Rat, zIm *big.Rat) color.Color {
// 	const iterations = 100
// 	const contrast = 15

// 	var vRe, vIm = new(big.Rat), new(big.Rat)

// 	// Reuse temporary variables to reduce allocations
// 	vRe2 := new(big.Rat)
// 	vIm2 := new(big.Rat)
// 	abs2 := new(big.Rat)
// 	vImNew := new(big.Rat)

// 	for n := uint8(0); n < iterations; n++ {
// 		// Compute vRe^2 and vIm^2
// 		vRe2.Mul(vRe, vRe)
// 		vIm2.Mul(vIm, vIm)

// 		// Compute |v|^2 = vRe^2 + vIm^2
// 		abs2.Add(vRe2, vIm2)

// 		// Early exit if magnitude exceeds 2
// 		if abs2.Cmp(four) > 0 {
// 			return color.Gray{255 - contrast*n}
// 		}

// 		// Compute 2*vRe*vIm + zIm
// 		vImNew.Mul(vRe, vIm)
// 		vImNew.Mul(vImNew, two)
// 		vImNew.Add(vImNew, zIm)

// 		// Compute vRe^2 - vIm^2 + zRe
// 		vRe.Sub(vRe2, vIm2)
// 		vRe.Add(vRe, zRe)

// 		// Update vIm
// 		vIm.Set(vImNew)
// 	}
// 	return color.Black
// }

package bigrat

import (
	"image"
	"image/color"
	"math/big"
)

// Render generates an image of the mandelbrot set.
func Render(xmin, ymin, xmax, ymax float64, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var cr, ci big.Rat
	for py := 0; py < height; py++ {
		y := ((float64(py) / float64(height)) * (ymax - ymin)) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			cr.SetFloat64(x)
			ci.SetFloat64(y)
			img.Set(px, py, mandelbrot(&cr, &ci))
		}
	}
	return img
}

func mandelbrot(cr, ci *big.Rat) color.Color {
	// https://randomascii.wordpress.com/2011/08/13/faster-fractals-through-algebra/
	const iterations = 200
	const contrast = 15

	var zr, zi, zrzi, zrsqr, zisqr big.Rat
	for n := uint8(0); n < iterations; n++ {
		zrzi.Add(&zr, &zi)
		zi.Mul(&zrzi, &zrzi).Sub(&zi, &zrsqr).Sub(&zi, &zisqr).Add(&zi, ci)
		zr.Sub(&zrsqr, &zisqr).Add(&zr, cr)
		zrsqr.Mul(&zr, &zr)
		zisqr.Mul(&zi, &zi)
		if new(big.Rat).Add(&zrsqr, &zisqr).Cmp(big.NewRat(4, 1)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
