package main

import "math"

func simpson(n int) float64 {
	a := 0.0
	b := math.Pi
	h := (b - a) / float64(n)

	f := func(x float64) float64 {
		sinX := math.Sin(x)
		return 1.5 * sinX * sinX * sinX
	}

	result := f(a) + f(b)

	for i := 1; i <= n/2; i++ {
		result += 4 * f(a+float64(2*i-1)*h)
	}

	for i := 1; i <= n/2-1; i++ {
		result += 2 * f(a+float64(2*i)*h)
	}

	result *= h / 3

	return result
}