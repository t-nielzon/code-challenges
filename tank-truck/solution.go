package main

import (
	"math"
)

func tankvol(h, d, vt int) int {
	if h == 0 || d == 0 || vt == 0 {
		return 0
	}

	r := float64(d) / 2.0
	L := float64(vt) / (math.Pi * r * r)

	hf := float64(h)
	A := r*r*math.Acos((r-hf)/r) - (r-hf)*math.Sqrt(hf*(2*r-hf))

	volume := A * L
	return int(volume)
}