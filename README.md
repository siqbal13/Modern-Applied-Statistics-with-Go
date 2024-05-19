# Modern-Applied-Statistics-with-Go
Bootstrapping in R and Go

Overview

This project demonstrates the implementation of bootstrap sampling in R and Go. The purpose is to evaluate the feasibility of using Go for modern applied statistics, particularly for reducing cloud computing costs.

Contents

RBootStraping.R: R implementation of the bootstrapping method.
RBootStraping_test.R: Unit tests for the R implementation.
RBootStraping_benchmark.R: Benchmark tests for the R implementation.
RBootStraping_profiling.R: Profiling tests for the R implementation.
main.go: Go implementation of the bootstrapping method.
main_test.go: Unit tests and benchmarks for the Go implementation.
profile.go: Profiling tests for the Go implementation.
README.md: This file, providing an overview and detailed documentation.
Requirements

R (tested on R 4.1.0)
Go (tested on Go 1.16)
Implementation

R Implementation

Implemented using the boot package to perform bootstrap sampling and estimate the standard error of the median.

Go Implementation

Implemented using the gonum package to perform the same bootstrap sampling.

Performance Comparison

Memory Usage:
R: 30 MB
Go: 10 MB
Processing Time:
R: 0.1412401 secs.
Go: 66.2249 msecs.
Recommendations

The transition from R to Go for computationally intensive tasks like bootstrapping can offer significant performance improvements and cost savings. While R remains a powerful tool for statistical analysis, Go provides a compelling alternative for scenarios requiring high performance and scalability. When to Use Go Over R: Large Datasets: Go's performance advantage becomes more significant with larger datasets due to its efficient use of multi-core processors. Cost Efficiency: Reduced execution time and memory usage translate to lower cloud computing costs. Scalability: Go is well-suited for high-performance, scalable applications

Cloud Provider

Provider: GCP (Google Cloud )
Machine Type: n1-standard-1.
Cost:
R: $0.000001 (approximately)
Go: $0.0000005 (approximately)
Estimated Savings: We recommend using Google Cloud Platform (GCP) for infrastructure as a service (IaaS). Based on our benchmarks, transitioning from R to Go could potentially save up to 50% in cloud computing costs due to reduced execution time and memory usage.
Usage

R

# Run the R script to perform bootstrap sampling
Rscript bootstrap.R
