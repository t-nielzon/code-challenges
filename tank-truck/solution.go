package kata

import "math"

func Tankvol(h, d, vt int) int {
	r := float64(d) / 2.0
	hh := float64(h)
	area := r*r*math.Acos((r-hh)/r) - (r-hh)*math.Sqrt(2*r*hh-hh*hh)
	ratio := area / (math.Pi * r * r)
	return int(math.Floor(ratio * float64(vt)))
}