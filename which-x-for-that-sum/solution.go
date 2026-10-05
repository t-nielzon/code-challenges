package kata

import "math"

func Solve(m float64) float64 {
	// the infinite series x + 2x^2 + 3x^3 + ... converges to x/(1-x)^2
	// given m = x/(1-x)^2, solve the quadratic mx^2 - (2m+1)x + m = 0
	// take the root that keeps x in (0, 1): ((2m+1) - sqrt(4m+1)) / (2m)
	return ((2*m + 1) - math.Sqrt(4*m+1)) / (2 * m)
}