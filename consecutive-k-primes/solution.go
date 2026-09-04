package main

func consec_kprimes(k int, arr []int) int {
	countPrimeFactors := func(n int) int {
		count := 0
		// Check for factor 2
		for n%2 == 0 {
			count++
			n /= 2
		}
		// Check for odd factors from 3 onwards
		for i := 3; i*i <= n; i += 2 {
			for n%i == 0 {
				count++
				n /= i
			}
		}
		// If n is still greater than 1, it's a prime factor
		if n > 1 {
			count++
		}
		return count
	}

	result := 0
	for i := 0; i < len(arr)-1; i++ {
		if countPrimeFactors(arr[i]) == k && countPrimeFactors(arr[i+1]) == k {
			result++
		}
	}
	return result
}