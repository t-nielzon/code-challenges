package kata

import "math"

func Fortune(f0 int, p float64, c0 float64, n int, i float64) bool {
	f := float64(f0)
	c := c0
	pRate := p / 100.0
	iRate := i / 100.0

	for year := 2; year <= n; year++ {
		f = math.Floor(f + f*pRate - c)
		c = math.Floor(c + c*iRate)
		if f < 0 {
			return false
		}
	}
	return true
}