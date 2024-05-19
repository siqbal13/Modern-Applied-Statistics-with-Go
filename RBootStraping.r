# RBootStraping.R

# Install and load the boot package
install.packages("boot")
library(boot)

start_time <- Sys.time()

# Function to calculate the median
median_func <- function(data, indices) {
  return(median(data[indices]))
}

# Generate sample data
set.seed(123)
data <- rnorm(1000)  # Normally distributed data

# Perform bootstrap
bootstrap_results <- boot(data, statistic = median_func, R = 1000)

# Print results
print(bootstrap_results)

# Calculate standard error of the median
se_median <- sd(bootstrap_results$t)
cat("Standard Error of the Median:", se_median, "\n")

# Calculate time taken by the bootstrap
end_time <- Sys.time()
time_taken <- end_time - start_time
print(time_taken)
