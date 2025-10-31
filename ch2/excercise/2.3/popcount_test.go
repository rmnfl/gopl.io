package popcount

import "testing"

var funcs = []struct {
	name string
	f    func(uint64) int
}{
	{"Expression", PopCountExpression},
	{"Loop", PopCountLoop},
}

func TestPopcount(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{1, 1},
		{255, 8},
		{256, 1},
		{0xAAAAAAAA, 16},
		{0xFFFFFFFF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
	}

	for _, f := range funcs {
		for _, test := range tests {
			t.Run(f.name, func(t *testing.T) {
				if got := f.f(test.input); got != test.want {
					t.Errorf("PopCount%s(%v) = %v, want %v", f.name, test.input, got, test.want)
				}
			})
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for _, f := range funcs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f.f(0xFFFFFFFFFFFFFFFF) // max uint64
			}
		})
	}
}
