package main

import "math"

func solve(m float64) float64 {
	// The series U(n, x) = x + 2x^2 + 3x^3 + ... + nx^n converges to x/(1-x)^2
	// Given m = x/(1-x)^2, we solve for x:
	// m(1-x)^2 = x
	// mx^2 - (2m+1)x + m = 0
	// Using quadratic formula: x = [(2m+1) ± √(4m+1)] / (2m)
	// We take the minus root to stay in range (0, 1)
	
	numerator := (2*m + 1) - math.Sqrt(4*m+1)
	denominator := 2 * m
	return numerator / denominator
}