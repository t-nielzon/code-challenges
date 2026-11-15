package kata

import "math"

func C(k int64) int64 {
	// Check if k is a perfect square
	sqrtFloat := math.Sqrt(float64(k))
	sqrt := int64(math.Round(sqrtFloat))
	
	if sqrt*sqrt != k {
		return 0
	}
	
	n := sqrt
	
	// Factorize n and count divisors of n^3
	// If k = n^2, then k^3 = n^6, so x*y = n^3
	// The answer is the number of divisors of n^3
	factors := make(map[int64]int64)
	temp := n
	for d := int64(2); d*d <= temp; d++ {
		for temp%d == 0 {
			factors[d]++
			temp /= d
		}
	}
	if temp > 1 {
		factors[temp]++
	}
	
	// Count divisors of n^3: (3*a1 + 1) * (3*a2 + 1) * ...
	count := int64(1)
	for _, exp := range factors {
		count *= (3*exp + 1)
	}
	
	return count
}