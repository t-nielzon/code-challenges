package kata

import "math"

const g = 9.81

func Dist(v, mu float64) float64 {
	u := v / 3.6
	return u + u*u/(2*mu*g)
}

func Speed(d, mu float64) float64 {
	// quadratic in u (m/s): u^2/(2*mu*g) + u - d = 0
	a := 1.0 / (2 * mu * g)
	u := (-1 + math.Sqrt(1+4*a*d)) / (2 * a)
	return u * 3.6
}