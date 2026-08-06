package main

func fortune(f0 float64, p float64, c0 float64, n int, i float64) bool {
	f := f0
	c := c0
	
	// Convert percentages to decimal
	p /= 100
	i /= 100
	
	for year := 2; year <= n; year++ {
		f = float64(int64(f*(1+p) - c))
		c = float64(int64(c * (1 + i)))
		
		if f < 0 {
			return false
		}
	}
	
	return true
}