package kata

func F(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	temp := n

	// Find prime factors and their exponents
	for p := 2; p*p <= temp; p++ {
		if temp%p == 0 {
			k := 0
			for temp%p == 0 {
				k++
				temp /= p
			}
			// Contribute k * p^(k-1) to result
			result *= k * pow(p, k-1)
		}
	}

	// If temp > 1, then it's a prime factor with exponent 1
	// In that case, contribute 1 * temp^0 = 1 (no change to result)

	return result
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}