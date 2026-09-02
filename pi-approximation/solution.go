package main

import "math"

func IterPi(epsilon float64) [2]interface{} {
	sum := 0.0
	n := 0
	pi := math.Pi
	
	for {
		sum += math.Pow(-1, float64(n)) / (2*float64(n) + 1)
		approximation := 4 * sum
		
		if math.Abs(approximation-pi) < epsilon {
			multiplier := math.Pow(10, 10)
			rounded := math.Round(approximation*multiplier) / multiplier
			return [2]interface{}{n + 1, rounded}
		}
		n++
	}
}