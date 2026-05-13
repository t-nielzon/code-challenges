package kata

import "math"

func tankvol(h, d, vt int) int {
	r := float64(d) / 2.0
	hf := float64(h)
	length := float64(vt) / (math.Pi * r * r)
	area := r*r*math.Acos((r-hf)/r) - (r-hf)*math.Sqrt(2*r*hf-hf*hf)
	return int(math.Floor(area * length))
}