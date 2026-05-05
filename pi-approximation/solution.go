package kata

import "math"

func IterPi(epsilon float64) [2]float64 {
	sum := 0.0
	i := 0
	for {
		i++
		sign := 1.0
		if i%2 == 0 {
			sign = -1.0
		}
		sum += sign / float64(2*i-1)
		pi := 4 * sum
		if math.Abs(pi-math.Pi) < epsilon {
			rounded := math.Round(pi*1e10) / 1e10
			return [2]float64{float64(i), rounded}
		}
	}
}