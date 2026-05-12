package kata

import "math"

func Tankvol(h, d, vt int) int {
	r := float64(d) / 2.0
	hf := float64(h)
	theta := 2 * math.Acos((r-hf)/r)
	area := 0.5 * r * r * (theta - math.Sin(theta))
	cylinderArea := math.Pi * r * r
	ratio := area / cylinderArea
	return int(math.Floor(ratio * float64(vt)))
}