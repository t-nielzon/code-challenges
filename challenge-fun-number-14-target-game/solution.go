package main

func maxScore(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	n := len(vals)
	// dp[i] = maximum points achievable starting from index i
	dp := make([]int, n+2)

	// Fill dp array from right to left
	for i := n - 1; i >= 0; i-- {
		// Option 1: skip shooting at i, take best from i+1
		skipCurrent := dp[i+1]
		// Option 2: shoot at i, skip i+1 due to reload, then best from i+2
		shootCurrent := vals[i] + dp[i+2]

		dp[i] = max(skipCurrent, shootCurrent)
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}