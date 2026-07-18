package main

func Phi(n int64) int64 {
	result := n

	// Check for factor 2
	if n%2 == 0 {
		for n%2 == 0 {
			n /= 2
		}
		result -= result / 2
	}

	// Check for odd factors from 3 onwards
	p := int64(3)
	for p*p <= n {
		if n%p == 0 {
			for n%p == 0 {
				n /= p
			}
			// Apply euler product formula: result * (p - 1) / p
			result -= result / p
		}
		p += 2
	}

	// If n > 1, then it's a prime factor
	if n > 1 {
		result -= result / n
	}

	return result
}