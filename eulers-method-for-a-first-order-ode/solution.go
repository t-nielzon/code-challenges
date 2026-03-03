package kata

import "math"

func ExEuler(n int) float64 {
	h := 1.0 / float64(n)
	y := 1.0
	sumErr := 0.0

	for k := 0; k <= n; k++ {
		x := float64(k) * h
		z := 1.0 + 0.5*math.Exp(-4.0*x) - 0.5*math.Exp(-2.0*x)
		if z != 0 {
			sumErr += math.Abs(y-z) / math.Abs(z)
		}
		// euler step for next iteration
		y += (2.0 - math.Exp(-4.0*x) - 2.0*y) * h
	}

	mean := sumErr / float64(n+1)
	return math.Floor(mean*1e6) / 1e6
}