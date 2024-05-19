# RBootStraping_profile.R

install.packages("profvis")
library(profvis)
install.packages("boot")
library(boot)

median_func <- function(data, indices) {
  return(median(data[indices]))
}

set.seed(123)
data <- rnorm(1000)

profvis({
  bootstrap_results <- boot(data, statistic = median_func, R = 1000)
  se_median <- sd(bootstrap_results$t)
  cat("Standard Error of the Median:", se_median, "\n")
})
