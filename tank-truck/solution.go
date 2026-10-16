package main

import "math"

func tankvol(h, d, vt int) int {
	if d == 0 || vt == 0 {
		return 0
	}

	r := float64(d) / 2.0
	h_float := float64(h)

	// area of circular segment using the formula:
	// A = r² * arccos((r - h) / r) - (r - h) * sqrt(2*r*h - h²)
	angle := math.Acos((r - h_float) / r)
	segment_area := r*r*angle - (r-h_float)*math.Sqrt(2*r*h_float-h_float*h_float)

	// length of cylinder from total volume
	cylinder_area := math.Pi * r * r
	length := float64(vt) / cylinder_area

	// volume of liquid
	volume := segment_area * length

	return int(math.Floor(volume))
}