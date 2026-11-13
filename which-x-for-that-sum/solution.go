package main

import "math"

func solve(m float64) float64 {
	// the infinite series U(∞, x) = Σ(n=1 to ∞) n*x^n = x/(1-x)²
	// given m = x/(1-x)², solve for x
	// let y = 1-x: m*y² = 1-y → m*y² + y - 1 = 0
	// quadratic formula: y = (-1 + √(1 + 4m)) / (2m)
	// therefore: x = 1 - y = (2m + 1 - √(1 + 4m)) / (2m)
	
	sqrtTerm := math.Sqrt(1 + 4*m)
	return (2*m + 1 - sqrtTerm) / (2 * m)
}