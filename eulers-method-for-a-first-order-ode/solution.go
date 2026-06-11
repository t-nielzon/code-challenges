package kata

import "math"

func ExEuler(n int) float64 {
	f := func(x, y float64) float64 {
		return 2 - math.Exp(-4*x) - 2*y
	}
	z := func(x float64) float64 {
		return 1 + 0.5*math.Exp(-4*x) - 0.5*math.Exp(-2*x)
	}

	h := 1.0 / float64(n)
	y := 1.0
	x := 0.0
	sum := 0.0

	for k := 0; k <= n; k++ {
		exact := z(x)
		sum += math.Abs(y-exact) / exact
		y += f(x, y) * h
		x = float64(k+1) * h
	}

	mean := sum / float64(n+1)
	return math.Trunc(mean*1e6) / 1e6
}