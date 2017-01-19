package main

import (
	"math/rand"
	"testing"
)

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
