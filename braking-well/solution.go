package kata

import "math"

const g = 9.81

func dist(v, mu float64) float64 {
	vMps := v / 3.6
	reaction := vMps * 1.0
	braking := vMps * vMps / (2 * mu * g)
	return reaction + braking
}

func speed(d, mu float64) float64 {
	// d = vMps + vMps^2/(2*mu*g) -> a*vMps^2 + vMps - d = 0
	a := 1.0 / (2 * mu * g)
	vMps := (-1.0 + math.Sqrt(1.0+4.0*a*d)) / (2.0 * a)
	return vMps * 3.6
}