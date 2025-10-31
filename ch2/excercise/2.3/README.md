# Exercise 2.3

Rewrite `popcount` to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows hwo to compare the performance of different implementations systematically.

## Benchmark Results

```
2.3 git:(master) ✗ go test -bench=.
goos: darwin
goarch: arm64
cpu: Apple M3 Pro
BenchmarkPopCount/Expression-12   614490846   1.920 ns/op
BenchmarkPopCount/Loop-12         208152464   5.761 ns/op
PASS
ok      gopl.io/ch2/excercise/2.3       3.375s
```

The expression-based version is about 3 times faster than the loop version.

