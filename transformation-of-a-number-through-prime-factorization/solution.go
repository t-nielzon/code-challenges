package main

func f(n int) int {
	if n <= 1 {
		return 1
	}

	// Find prime factorization
	primeFactors := make(map[int]int)
	temp := n

	for i := 2; i*i <= temp; i++ {
		for temp%i == 0 {
			primeFactors[i]++
			temp /= i
		}
	}
	if temp > 1 {
		primeFactors[temp]++
	}

	// Calculate geometric derivative
	result := 1
	for prime, exponent := range primeFactors {
		if exponent > 1 {
			power := 1
			for i := 0; i < exponent-1; i++ {
				power *= prime
			}
			result *= exponent * power
		}
	}

	return result
}