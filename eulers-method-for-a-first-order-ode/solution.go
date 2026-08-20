package main

import "math"

func ExEuler(n int) float64 {
	x := 0.0
	y := 1.0
	h := 1.0 / float64(n)

	f := func(x, y float64) float64 {
		return 2 - math.Exp(-4*x) - 2*y
	}

	z := func(x float64) float64 {
		return 1 + 0.5*math.Exp(-4*x) - 0.5*math.Exp(-2*x)
	}

	var totalError float64

	for k := 0; k <= n; k++ {
		zk := z(x)
		error := math.Abs(y-zk) / zk
		totalError += error

		if k < n {
			y = y + f(x, y)*h
			x = x + h
		}
	}

	meanError := totalError / float64(n+1)
	return math.Floor(meanError*1e6) / 1e6
}