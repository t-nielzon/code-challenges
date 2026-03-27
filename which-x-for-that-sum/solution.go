package kata

import "math"

func Solve(m float64) float64 {
	// sum k*x^k (k=1..inf) = x/(1-x)^2, so m = x/(1-x)^2
	// solving mx^2 - (2m+1)x + m = 0 via quadratic formula
	return (2*m + 1 - math.Sqrt(4*m+1)) / (2 * m)
}