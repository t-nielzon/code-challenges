package main

import "fmt"

func PrimeFactors(n int) string {
	result := ""

	addFactor := func(prime, count int) {
		if count == 1 {
			result += fmt.Sprintf("(%d)", prime)
		} else {
			result += fmt.Sprintf("(%d**%d)", prime, count)
		}
	}

	// check for factor 2
	count := 0
	for n%2 == 0 {
		count++
		n = n / 2
	}
	if count > 0 {
		addFactor(2, count)
	}

	// check for odd factors from 3 onwards
	for i := 3; i*i <= n; i += 2 {
		count = 0
		for n%i == 0 {
			count++
			n = n / i
		}
		if count > 0 {
			addFactor(i, count)
		}
	}

	// if n is still greater than 1, then it's a prime factor
	if n > 1 {
		addFactor(n, 1)
	}

	return result
}