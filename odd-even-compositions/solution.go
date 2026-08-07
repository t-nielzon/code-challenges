package main

func OddEvenCompositions(n int) int {
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			// Even: valid choices are 1, 2, 4, 6, ..., i
			dp[i] = dp[i-1] // choice of 1
			for k := 2; k <= i; k += 2 {
				dp[i] += dp[i-k]
			}
		} else {
			// Odd: valid choices are 1, 3, 5, 7, ..., i
			for k := 1; k <= i; k += 2 {
				dp[i] += dp[i-k]
			}
		}
	}

	return dp[n]
}