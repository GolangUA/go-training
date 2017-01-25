package main

import (
	"math"
	"math/rand"
	"testing"
)

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BenchmarkMin(b *testing.B) {
	a, c := rand.Int(), rand.Int()
	for i := 0; i < b.N; i++ {
		min(a, c)
	}
}

func BenchmarkMin2(b *testing.B) {
	a, c := rand.Int(), rand.Int()
	for i := 0; i < b.N; i++ {
		min2(a, c)
	}
}
