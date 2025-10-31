# Exercise 2.5

The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of `popcount` that counts bits by using this fact, and asses its performance.

## Benchmark Results

```
➜  2.5 git:(master) ✗ go test -bench=.
goos: darwin
goarch: arm64
pkg: gopl.io/ch2/excercise/2.5
cpu: Apple M3 Pro
BenchmarkPopCount/Expression-12                 609194520                1.921 ns/op
BenchmarkPopCount/Clear-12                      35681275                33.41 ns/op
PASS
ok      gopl.io/ch2/excercise/2.5       2.803s
```

The table-lookup (expression) version is much faster than the clear-rightmost-bit version (x&(x-1)). On this machine the clear-based implementation is about 17x slower than the expression-based one.

