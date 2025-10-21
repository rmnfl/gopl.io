package echo

import (
	"strconv"
	"testing"
)

func makeArgs(count int) []string {
	args := make([]string, count)

	for i := 0; i < count; i++ {
		args[i] = strconv.Itoa(i)
	}

	return args
}

func BenchmarkEcho1_100(b *testing.B)   { benchmarkEcho(b, echo1, 100) }
func BenchmarkEcho1_1000(b *testing.B)  { benchmarkEcho(b, echo1, 1000) }
func BenchmarkEcho1_10000(b *testing.B) { benchmarkEcho(b, echo1, 10000) }

func BenchmarkEcho2_100(b *testing.B)   { benchmarkEcho(b, echo2, 100) }
func BenchmarkEcho2_1000(b *testing.B)  { benchmarkEcho(b, echo2, 1000) }
func BenchmarkEcho2_10000(b *testing.B) { benchmarkEcho(b, echo2, 10000) }

func BenchmarkEcho3_100(b *testing.B)   { benchmarkEcho(b, echo3, 100) }
func BenchmarkEcho3_1000(b *testing.B)  { benchmarkEcho(b, echo3, 1000) }
func BenchmarkEcho3_10000(b *testing.B) { benchmarkEcho(b, echo3, 10000) }

func benchmarkEcho(b *testing.B, fn func([]string) string, count int) {
	args := makeArgs(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fn(args)
	}
}
