package kata

import "math/big"

func OddEvenCompositions(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	// dp[r] = number of valid odd-even compositions for remainder r
	// allowed parts for remainder r:
	//   - 1 (always allowed), reduces remainder by 1
	//   - r itself (always allowed), terminates the composition
	//   - any k (2 <= k < r) where k and r share the same parity
	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(1) // base: empty remainder means we completed a composition

	for r := 1; r <= n; r++ {
		dp[r] = new(big.Int)
		// part = r (the number itself): always valid, contributes 1 composition
		dp[r].Add(dp[r], dp[0]) // equivalent to adding 1

		// part = 1: always valid, remainder becomes r-1
		if r >= 1 {
			dp[r].Add(dp[r], dp[r-1])
		}

		// parts k where 2 <= k < r and k shares parity with r
		for k := 2; k < r; k++ {
			if k%2 == r%2 {
				dp[r].Add(dp[r], dp[r-k])
			}
		}
	}

	return dp[n]
}