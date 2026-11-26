package main

import "math"

func GiveTriang(maxPerimeter int) int {
	count := 0
	
	for a := 1; a < maxPerimeter; a++ {
		for b := a; b < maxPerimeter; b++ {
			// law of cosines: c² = a² + b² - 2ab·cos(120°)
			// since cos(120°) = -1/2: c² = a² + b² + ab
			c_squared := a*a + b*b + a*b
			c := int(math.Sqrt(float64(c_squared)) + 0.5)
			
			if c*c == c_squared {
				if a + b + c <= maxPerimeter {
					count++
				} else {
					break
				}
			}
		}
	}
	
	return count
}