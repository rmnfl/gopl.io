# Exercise 2.4

Write a version of `popcount` that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time.  Compare its performance to the table-lookup version.

## Benchmark Results

```
➜  2.4 git:(master) ✗ go test -bench=.
goos: darwin
goarch: arm64
cpu: Apple M3 Pro
BenchmarkPopCount/Expression-12   614641792   1.957 ns/op
BenchmarkPopCount/Shift-12        42682818   28.28 ns/op
PASS
ok      gopl.io/ch2/excercise/2.4       3.852s
```

The table-lookup version is much faster than the bit-shifting version.

