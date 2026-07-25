package kata

import "math"

func C(k int64) int {
	// For z = k in z³ = x² * y², we have k³ = x² * y²
	// This means (xy)² = k³
	// For integer solutions, k must be a perfect square: k = m²
	// Then k³ = m⁶ = (m³)², so xy = m³
	// Count the number of positive divisor pairs (x,y) where xy = m³
	// This equals the number of divisors of m³

	// Check if k is a perfect square
	m := int64(math.Sqrt(float64(k)))

	// Handle floating point precision
	if m*m != k {
		if (m+1)*(m+1) == k {
			m++
		} else {
			return 0
		}
	}

	// Factorize m
	factors := make(map[int64]int)
	temp := m

	for d := int64(2); d*d <= temp; d++ {
		for temp%d == 0 {
			factors[d]++
			temp /= d
		}
	}
	if temp > 1 {
		factors[temp]++
	}

	// Count divisors of m³
	// If m = p₁^a₁ * p₂^a₂ * ... then m³ = p₁^(3a₁) * p₂^(3a₂) * ...
	// Number of divisors = (3a₁+1)(3a₂+1)...(3aₙ+1)
	count := 1
	for _, exp := range factors {
		count *= (3*exp + 1)
	}

	return count
}