package kata

import "math"

func Iter_pi(epsilon float64) [2]float64 {
	sum := 0.0
	iterations := 0
	sign := 1.0

	for math.Abs(math.Pi-4*sum) >= epsilon {
		sum += sign / float64(2*iterations+1)
		sign = -sign
		iterations++
	}

	approx := math.Round(4*sum*1e10) / 1e10
	return [2]float64{float64(iterations), approx}
}