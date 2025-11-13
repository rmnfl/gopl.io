package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type config struct {
	width, height       int
	cells               int
	xyrange             float64
	xyscale, zscale     float64
	angle, sin30, cos30 float64
	color               string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	cfg := defaultConfig()
	cfg = applyQueryParams(cfg, r.URL.Query())

	surface(w, cfg)
}

func defaultConfig() config {
	return config{
		width:   600,
		height:  320,
		cells:   100,
		xyrange: 30.0,
		angle:   math.Pi / 6,
		color:   "#ffffff",
	}
}

func applyQueryParams(cfg config, q url.Values) config {
	if v := parseInt(q.Get("width")); v > 0 {
		cfg.width = v
	}
	if v := parseInt(q.Get("height")); v > 0 {
		cfg.height = v
	}
	if s := q.Get("color"); s != "" {
		cfg.color = s
	}
	cfg.xyscale = float64(cfg.width) / 2 / cfg.xyrange
	cfg.zscale = float64(cfg.height) * 0.4
	cfg.sin30, cfg.cos30 = math.Sin(cfg.angle), math.Cos(cfg.angle)
	return cfg
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func surface(out io.Writer, cfg config) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", cfg.width, cfg.height)
	for i := 0; i < cfg.cells; i++ {
		for j := 0; j < cfg.cells; j++ {
			ax, ay, aok := corner(i+1, j, cfg)
			bx, by, bok := corner(i, j, cfg)
			cx, cy, cok := corner(i, j+1, cfg)
			dx, dy, dok := corner(i+1, j+1, cfg)

			if !(aok && bok && cok && dok) {
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, cfg.color)
		}
	}
	fmt.Fprint(out, "</svg>")
}

func corner(i, j int, cfg config) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := cfg.xyrange * (float64(i)/float64(cfg.cells) - 0.5)
	y := cfg.xyrange * (float64(j)/float64(cfg.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(cfg.width)/2 + (x-y)*cfg.cos30*cfg.xyscale
	sy := float64(cfg.height)/2 + (x+y)*cfg.sin30*cfg.xyscale - z*cfg.zscale

	if math.IsNaN(sx) || math.IsInf(sx, 0) || math.IsNaN(sy) || math.IsInf(sy, 0) {
		return 0, 0, false
	}
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
