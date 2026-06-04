package kata

import "math"

// rationalize sqrt(1+x)-1 by multiplying by (sqrt(1+x)+1) to avoid the
// catastrophic cancellation of subtracting two nearly equal numbers
func F(x float64) float64 {
	return x / (math.Sqrt(1+x) + 1)
}