package main

import (
	"math"
)

func Ex_euler(n int) float64 {
	h := 1.0 / float64(n)
	x := 0.0
	y := 1.0
	
	totalError := 0.0
	
	for k := 0; k <= n; k++ {
		// exact solution at x_k
		z := 1.0 + 0.5*math.Exp(-4*x) - 0.5*math.Exp(-2*x)
		
		// relative error
		relError := math.Abs(y-z) / z
		totalError += relError
		
		// euler's method step: f(x,y) = 2 - e^(-4x) - 2y
		f := 2.0 - math.Exp(-4*x) - 2*y
		y = y + f*h
		x = x + h
	}
	
	// mean error
	meanError := totalError / float64(n+1)
	
	// truncate to 6 decimal places
	return math.Trunc(meanError*1e6) / 1e6
}