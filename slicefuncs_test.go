package slicefuncs

import (
	"math/rand"
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	s := make([]int, 3000)
	for i := 0; i < 3000; i++ {
		s[i] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		Filter(s, func(n int) bool { return n%2 == 0 })
	}
}

func BenchmarkFlatten(b *testing.B) {
	s := make([][]int, 20)

	for i := 0; i < 20; i++ {
		l := rand.Intn(10)
		m := make([]int, l)
		for j := 0; j < l; j++ {
			m[j] = rand.Int()
		}
		s[i] = m
	}
	for i := 0; i < b.N; i++ {
		Flatten(s)
	}
}
