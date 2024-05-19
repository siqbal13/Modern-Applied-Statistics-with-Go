package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Bootstrap function to calculate the median
func bootstrap(data []float64, n int, seed int64) []float64 {
	rand.Seed(seed)
	medians := make([]float64, n)
	for i := 0; i < n; i++ {
		sample := sampleWithReplacement(data, len(data))
		medians[i] = median(sample)
	}
	return medians
}

// Sample with replacement
func sampleWithReplacement(data []float64, size int) []float64 {
	sample := make([]float64, size)
	for i := range sample {
		sample[i] = data[rand.Intn(len(data))]
	}
	return sample
}

// Calculate the median
func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}


func main() {
	// Generate sample data
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, 1000)
	for i := range data {
		data[i] = rand.NormFloat64()
	}

	// Perform bootstrap
	start := time.Now()
	bootstrapResults := bootstrap(data, 1000, time.Now().UnixNano())
	elapsed := time.Since(start)
	fmt.Printf("Bootstrapping completed in %s\n", elapsed)

	// Calculate standard error of the median
	sum := 0.0
	for _, v := range bootstrapResults {
		sum += v
	}
	mean := sum / float64(len(bootstrapResults))

	sumSquares := 0.0
	for _, v := range bootstrapResults {
		sumSquares += (v - mean) * (v - mean)
	}
	se := sumSquares / float64(len(bootstrapResults))
	fmt.Printf("Standard Error of the Median: %f\n", se)
}
