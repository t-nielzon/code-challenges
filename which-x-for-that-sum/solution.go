package solution

import "math"

func Solve(m float64) float64 {
	// the infinite series sum k*x^k from k=1 to infinity equals x/(1-x)^2
	// given m, we solve: m = x/(1-x)^2
	// let y = 1-x, then: m*y^2 + y - 1 = 0
	// using quadratic formula: y = (-1 + sqrt(1 + 4m)) / (2m)
	// therefore: x = 1 - y
	
	discriminant := 1.0 + 4.0*m
	y := (-1.0 + math.Sqrt(discriminant)) / (2.0 * m)
	return 1.0 - y
}