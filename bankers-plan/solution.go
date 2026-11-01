// solution.go
package main

import "math"

func fortune(f0 int, p int, c0 int, n int, i int) bool {
	f := float64(f0)
	c := float64(c0)
	pRate := float64(p) / 100.0
	iRate := float64(i) / 100.0

	for year := 1; year < n; year++ {
		// apply interest and withdraw
		f = f + f*pRate - c
		f = math.Floor(f)

		if f < 0 {
			return false
		}

		// apply inflation to withdrawal amount for next year
		c = c + c*iRate
		c = math.Floor(c)
	}

	return true
}