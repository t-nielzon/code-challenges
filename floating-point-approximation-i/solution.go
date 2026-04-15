package kata

import "math"

func F(x float64) float64 {
	return x / (math.Sqrt(1+x) + 1)
}