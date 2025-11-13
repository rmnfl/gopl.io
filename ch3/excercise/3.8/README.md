# Exercise 3.8: Mandelbrot Fractal with Multiple Number Representations

Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter two types are found in the math/big package. Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?

## Usage

Each numeric representation is in its own package. Run benchmarks to compare performance:

```bash
# Run all benchmarks (excluding bigrat which times out)
go test -bench=Benchmark -benchmem -run=^$

# Run a specific benchmark
go test -bench=BenchmarkComplex64 -benchmem -run=^$
go test -bench=BenchmarkComplex128 -benchmem -run=^$
go test -bench=BenchmarkBigFloat -benchmem -run=^$
```

To render zoom-level test images (saves PNG files to `images/` directory):

```bash
go test -v -run TestZoomLevels
```

## Benchmark Results

```bash
➜  3.8 git:(master) ✗ go test -bench=Benchmark -benchmem -run=^$ 
goos: darwin
goarch: arm64
pkg: gopl.io/ch3/excercise/3.8
cpu: Apple M3 Pro
BenchmarkComplex64-12                392           3031101 ns/op          323908 B/op      40002 allocs/op
BenchmarkComplex128-12               524           2273934 ns/op          323906 B/op      40002 allocs/op
BenchmarkBigFloat-12                   7         147422589 ns/op        141474368 B/op   4019166 allocs/op
PASS
ok      gopl.io/ch3/excercise/3.8       4.245s
```
### Why BigRat Is Disabled

The `bigrat` package implementation is **commented out and not included in benchmarks** because it is too slow:

- A single 200×200 image render with bigrat exceeds the 10-minute test timeout
- Each pixel requires:
  - `SetFloat64()` conversions (expensive rational parsing)
  - Per-iteration allocations of temporary big.Rat objects
  - Heavy garbage collection pressure from millions of small heap allocations
- Rational arithmetic overhead is simply incompatible with pixel-by-pixel iteration over millions of points

**Conclusion:** Unbounded-precision rationals are mathematically elegant but computationally infeasible for iterative fractal rendering without dramatic architectural changes (e.g., lazy evaluation, memoization, or algorithm restructuring).


## Analysis

### Precision vs. Zoom Level

Empirical observation from rendering test images at zoom levels 1e0 to 1e20:

| Implementation | Artifacts Start | Notes |
|---|---|---|
| **complex64** | Zoom ≥ 1e6 | Single-precision float32 exhausts precision very quickly |
| **complex128** | Zoom ≥ 1e15 | Double-precision float64 handles up to ~15 significant digits |
| **bigfloat** | Zoom ≥ 1e15 | Same precision limit as complex128 (bounded arbitrary-precision) |
| **bigrat** | Never tested | Too slow to render (disabled) |

### Performance Trade-offs

- **complex64**: ~3ms per 200×200 image, but only usable for zoom < 1e6
- **complex128**: ~2.3ms per 200×200 image, usable up to zoom 1e15
- **bigfloat**: ~147ms per 200×200 image (~65x slower), but identical precision to complex128 with this default precision
- **bigrat**: Timeouts (>600s for 200×200 image)