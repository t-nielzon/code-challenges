package solution

func KprimesStep(k, step, start, nd int) [][]int {
	// count prime factors with multiplicity
	countPrimeFactors := func(n int) int {
		if n <= 1 {
			return 0
		}
		count := 0

		// divide by 2
		for n%2 == 0 {
			count++
			n /= 2
		}

		// check odd divisors
		for i := 3; i*i <= n; i += 2 {
			for n%i == 0 {
				count++
				n /= i
			}
		}

		// if n > 1, it's a prime factor
		if n > 1 {
			count++
		}

		return count
	}

	var result [][]int

	// find all k-primes and check if next k-prime is at step distance
	for i := start; i <= nd; i++ {
		if countPrimeFactors(i) == k && i+step <= nd && countPrimeFactors(i+step) == k {
			result = append(result, []int{i, i + step})
		}
	}

	return result
}