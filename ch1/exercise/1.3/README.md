# Exercise 1.3

Experiment to measure the difference in running time between our potentially inefficient versions and
the one that uses `strings.Join`.  
(Section 1.6 illustrates part of the `time` package, and Section 11.4 shows how to write benchmark tests
for systematic performance evaluation.)

---

## Implementation

Three versions of the `echo` program were benchmarked:

- **Echo1** – concatenates strings with `+=` using an index loop  
- **Echo2** – concatenates strings with `+=` using a `range` loop  
- **Echo3** – uses the efficient `strings.Join`

All three functions take generated arguments (100, 1000, 10000 elements) to simulate different input sizes.

---

## Results

### Benchmarking
```bash
➜  1.3 git:(master) ✗ go test -bench=.
goos: darwin
goarch: arm64
pkg: gopl.io/ch1/exercise/1.3
cpu: Apple M3 Pro
BenchmarkEcho1_100-12             410536              2924 ns/op
BenchmarkEcho1_1000-12              6885            171583 ns/op
BenchmarkEcho1_10000-12               57          20780514 ns/op
BenchmarkEcho2_100-12             393240              2959 ns/op
BenchmarkEcho2_1000-12              6901            172299 ns/op
BenchmarkEcho2_10000-12               57          20615608 ns/op
BenchmarkEcho3_100-12            1916725               630.1 ns/op
BenchmarkEcho3_1000-12            170641              6728 ns/op
BenchmarkEcho3_10000-12            17227             69855 ns/op
PASS
ok      gopl.io/ch1/exercise/1.3        13.514s
```

## Analysis

| Function  | Args = 100 | Args = 1 000 | Args = 10 000 | Growth Pattern |
| --------- | ---------- | ------------ | ------------- | -------------- |
| **Echo1** | 2.9 µs/op  | 171 µs/op    | 20.8 ms/op    | ~O(n²)         |
| **Echo2** | 2.9 µs/op  | 172 µs/op    | 20.6 ms/op    | ~O(n²)         |
| **Echo3** | 0.63 µs/op | 6.7 µs/op    | 69 µs/op      | ~O(n)          |

- Both Echo1 and Echo2 show quadratic time growth, caused by repeated string concatenation.
- Echo3 using strings.Join exhibits linear performance, scaling efficiently even for large inputs.

## Conclusion

The benchmark confirms the theoretical expectation:
Naive concatenation `(+=)` becomes extremely slow for large argument lists due to repeated memory allocation and copying.
`strings.Join` is orders of magnitude faster and scales linearly with input size.

### Takeaway:
Always use strings.Join for concatenating many strings in Go — it’s both faster and more memory-efficient.