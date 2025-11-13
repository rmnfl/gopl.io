package mandelbrot

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopl.io/ch3/excercise/3.8/bigfloat"
	"gopl.io/ch3/excercise/3.8/complex128"
	"gopl.io/ch3/excercise/3.8/complex64"
)

// Small image for quick benchmarks
const width, height = 200, 200
const xmin, ymin, xmax, ymax = -2, -2, 2, 2

func BenchmarkComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex64.Render(xmin, xmax, ymin, ymax, width, height)
	}
}

func BenchmarkComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex128.Render(xmin, xmax, ymin, ymax, width, height)
	}
}

func BenchmarkBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigfloat.Render(xmin, xmax, ymin, ymax, width, height)
	}
}

// func BenchmarkBigRat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		bigrat.Render(xmin, xmax, ymin, ymax, width, height)
// 	}
// }

// TestZoomLevels renders mandelbrot sets at increasing zoom levels
// and saves them as PNG files to compare artifact appearance across
// different numeric representations.
func TestZoomLevels(t *testing.T) {
	outputDir := "images"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	// Center of interesting region in mandelbrot set
	centerX, centerY := -1.8, 0.0

	// Define zoom levels to test (from no zoom to high zoom).
	// Build zoomFactors from a range of exponents so it's easy to change.
	zoomFactors := make([]float64, 0)
	for e := 0; e <= 20; e += 1 {
		zoomFactors = append(zoomFactors, math.Pow10(e))
	}

	implementations := map[string]func(float64, float64, float64, float64, int, int) image.Image{
		"complex64":  complex64.Render,
		"complex128": complex128.Render,
		"bigfloat":   bigfloat.Render,
		// "bigrat":     bigrat.Render,
	}

	for name, renderFunc := range implementations {
		implDir := filepath.Join(outputDir, name)
		if err := os.MkdirAll(implDir, 0755); err != nil {
			t.Fatalf("Failed to create implementation directory: %v", err)
		}

		for _, zoom := range zoomFactors {
			// Calculate new bounds with zoom
			halfWidth := 1.0 / zoom  // Initial width is 2 (-1 to 1)
			halfHeight := 1.0 / zoom // Initial height is 2 (-1 to 1)

			xminZ := centerX - halfWidth
			xmaxZ := centerX + halfWidth
			yminZ := centerY - halfHeight
			ymaxZ := centerY + halfHeight

			// Render image
			img := renderFunc(xminZ, xmaxZ, yminZ, ymaxZ, width, height)

			// Save as PNG. Use scientific (1e) notation for zoom level in filenames.
			zoomLabel := fmt.Sprintf("%.0e", zoom)
			// Remove any '+' signs so names are like 1e10 instead of 1e+10
			zoomLabel = strings.ReplaceAll(zoomLabel, "+", "")
			filename := filepath.Join(implDir, fmt.Sprintf("zoom_%s.png", zoomLabel))
			f, err := os.Create(filename)
			if err != nil {
				t.Fatalf("Failed to create file %s: %v", filename, err)
			}
			defer f.Close()

			if err := png.Encode(f, img); err != nil {
				t.Fatalf("Failed to encode PNG: %v", err)
			}

			t.Logf("Saved %s - zoom %sx", filename, zoomLabel)
		}
	}

	t.Logf("All zoom level images saved to %s/", outputDir)
}
