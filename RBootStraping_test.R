# RBootStraping_test.R

install.packages("testthat")
library(testthat)
install.packages("boot")
library(boot)

median_func <- function(data, indices) {
  return(median(data[indices]))
}

test_that("median_func calculates the correct median", {
  data <- c(1, 3, 5, 7, 9)
  indices <- c(1, 2, 3, 4, 5)
  expect_equal(median_func(data, indices), 5)
})

test_that("bootstrap results have the correct length", {
  set.seed(123)
  data <- rnorm(1000)
  bootstrap_results <- boot(data, statistic = median_func, R = 1000)
  expect_equal(length(bootstrap_results$t), 1000)
})
