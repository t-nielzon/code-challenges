package kata

import "math"

func tankvol(h, d, vt int) int {
	r := float64(d) / 2.0
	hf := float64(h)
	// area of the circular segment filled by the liquid in the tank's vertical cross-section
	segment := r*r*math.Acos((r-hf)/r) - (r-hf)*math.Sqrt(2*r*hf-hf*hf)
	// liquid volume is the segment's fraction of the full circle times total volume
	fraction := segment / (math.Pi * r * r)
	return int(fraction * float64(vt))
}