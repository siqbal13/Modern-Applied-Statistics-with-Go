package main

import (
	"testing"
)

func TestMedian(t *testing.T) {
	data := []float64{1, 3, 5, 7, 9}
	expected := 5.0
	result := median(data)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestBootstrapLength(t *testing.T) {
	data := make([]float64, 1000)
	for i := range data {
		data[i] = float64(i)
	}
	results := bootstrap(data, 1000, 1)
	if len(results) != 1000 {
		t.Errorf("Expected length %d, got %d", 1000, len(results))
	}
}

func BenchmarkBootstrap(b *testing.B) {
	data := make([]float64, 1000)
	for i := range data {
		data[i] = float64(i)
	}
	for i := 0; i < b.N; i++ {
		bootstrap(data, 1000, 1)
	}
}
