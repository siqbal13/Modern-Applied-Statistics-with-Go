# RBootStraping_benchmark.R

install.packages("microbenchmark")
library(microbenchmark)
install.packages("boot")
library(boot)

median_func <- function(data, indices) {
  return(median(data[indices]))
}

set.seed(123)
data <- rnorm(1000)

benchmark_results <- microbenchmark(
  boot(data, statistic = median_func, R = 1000),
  times = 10
)
print(benchmark_results)
