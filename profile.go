package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile: ", err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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

	// Memory profiling
	memf, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create memory profile: ", err)
		return
	}
	defer memf.Close()
	pprof.WriteHeapProfile(memf)
}
