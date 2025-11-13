package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

type config struct {
	x    float64
	y    float64
	zoom float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cfg := defaultConfig()
	cfg = applyQueryParams(cfg, r.URL.Query())

	renderImage(cfg, w)
}

func defaultConfig() config {
	return config{
		x:    0,
		y:    0,
		zoom: 1,
	}
}

func applyQueryParams(cfg config, q url.Values) config {
	if v, err := strconv.ParseFloat(q.Get("x"), 64); err == nil {
		cfg.x = v
	}
	if v, err := strconv.ParseFloat(q.Get("y"), 64); err == nil {
		cfg.y = v
	}
	if v, err := strconv.ParseFloat(q.Get("zoom"), 64); err == nil {
		cfg.zoom = v
	}
	return cfg
}

func renderImage(cfg config, w http.ResponseWriter) {
	const (
		width, height = 1024, 1024
	)

	scale := 1 / cfg.zoom
	xmin, xmax := cfg.x-2*scale, cfg.x+2*scale
	ymin, ymax := cfg.y-2*scale, cfg.y+2*scale

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
	png.Encode(w, img) // NOTE: ignoring errors
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
