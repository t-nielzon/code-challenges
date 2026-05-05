package kata

import "math"

func Going(n int) float64 {
	sum := 0.0
	term := 1.0
	for k := n; k >= 1; k-- {
		sum += term
		term /= float64(k)
	}
	return math.Floor(sum*1e6) / 1e6
}