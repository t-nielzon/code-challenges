func consec_kprimes(k int, arr []int) int {
	countPrimeFactors := func(n int) int {
		count := 0
		// factor out all 2s
		for n%2 == 0 {
			count++
			n /= 2
		}
		// factor out odd numbers
		for i := 3; i*i <= n; i += 2 {
			for n%i == 0 {
				count++
				n /= i
			}
		}
		// if n is still greater than 1, it's a prime factor itself
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