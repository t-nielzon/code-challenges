// solution.go
package kata

import (
	"math"
)

func interp(f func(float64) float64, l, u float64, n int) []float64 {
	d := (u - l) / float64(n)
	result := make([]float64, n)
	
	for i := 0; i < n; i++ {
		x := l + float64(i)*d
		y := f(x)
		result[i] = math.Floor(y*100.0) / 100.0
	}
	
	return result
}