package main

func ColorChoice(m int, n int) int {
	// Helper function to calculate binomial coefficient C(n, k)
	binomial := func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		if k == 0 || k == n {
			return 1
		}
		// Use symmetry to reduce computation
		if k > n-k {
			k = n - k
		}
		result := int64(1)
		for i := 0; i < k; i++ {
			result = result * int64(n-i) / int64(i+1)
		}
		return result
	}

	target := int64(m)

	// Find x such that C(n, x) = m
	for x := 0; x <= n; x++ {
		if binomial(n, x) == target {
			return x
		}
	}

	return -1
}