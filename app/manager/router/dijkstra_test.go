package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dijkstraRouteAlgorithm(t *testing.T) {
	setupTestData(t)

	dijkstraRouteAlgorithm()

	assert.Equal(t, decisions["B"], nodes["B"])
	assert.Equal(t, decisions["C"], nodes["B"])
	assert.Equal(t, decisions["D"], nodes["B"])
	assert.Equal(t, decisions["E"], nodes["E"])
}

func Benchmark_dijkstraRouteAlgorithm(b *testing.B) {
	b.Run("dijkstraRouteAlgorithm 1000", func(b *testing.B) {
		setupBenchmarkData(b, 1000)
		for i := 0; i < b.N; i++ {
			dijkstraRouteAlgorithm()
		}
	})

	b.Run("dijkstraRouteAlgorithm 10000", func(b *testing.B) {
		setupBenchmarkData(b, 10000)
		for i := 0; i < b.N; i++ {
			dijkstraRouteAlgorithm()
		}
	})
}
