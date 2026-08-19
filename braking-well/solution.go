package main

import "math"

func dist(v, mu float64) float64 {
	vms := v / 3.6
	return vms + (vms*vms)/(2*mu*9.81)
}

func speed(d, mu float64) float64 {
	coeff := 2 * mu * 9.81
	discriminant := coeff*coeff + 4*d*coeff
	vms := (-coeff + math.Sqrt(discriminant)) / 2
	return vms * 3.6
}