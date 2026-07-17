func TargetGame(vals []int) int {
	n := len(vals)
	if n == 0 {
		return 0
	}

	if n == 1 {
		if vals[0] < 0 {
			return 0
		}
		return vals[0]
	}

	dp := make([]int, n)

	dp[0] = vals[0]
	if dp[0] < 0 {
		dp[0] = 0
	}

	dp[1] = vals[1]
	if dp[1] < dp[0] {
		dp[1] = dp[0]
	}

	for i := 2; i < n; i++ {
		option1 := dp[i-1]
		option2 := vals[i] + dp[i-2]
		dp[i] = option1
		if option2 > option1 {
			dp[i] = option2
		}
	}

	return dp[n-1]
}