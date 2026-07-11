package main

func primeFactorize(n int64) map[int64]int64 {
	factors := make(map[int64]int64)

	for n%2 == 0 {
		factors[2]++
		n /= 2
	}

	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	if n > 1 {
		factors[n]++
	}

	return factors
}

func PrimeOperations(x, y int64) int {
	xFactors := primeFactorize(x)
	yFactors := primeFactorize(y)

	allPrimes := make(map[int64]bool)
	for p := range xFactors {
		allPrimes[p] = true
	}
	for p := range yFactors {
		allPrimes[p] = true
	}

	operations := 0
	for p := range allPrimes {
		xPower := xFactors[p]
		yPower := yFactors[p]
		diff := xPower - yPower
		if diff < 0 {
			diff = -diff
		}
		operations += int(diff)
	}

	return operations
}