package tests

import (
	"testing"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////
//go test -run=Negative -bench=.

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
