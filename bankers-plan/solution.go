package main

import "math"

func fortune(f0 float64, p float64, c0 float64, n int, i float64) bool {
	f := f0
	c := c0

	// Convert percentages to decimal form
	p = p / 100
	i = i / 100

	// Simulate n-1 years
	for year := 1; year < n; year++ {
		// Apply interest and withdraw
		f = math.Floor(f*(1+p) - c)

		// If balance is negative, plan fails
		if f < 0 {
			return false
		}

		// Adjust withdrawal for inflation
		c = math.Floor(c * (1 + i))
	}

	return true
}