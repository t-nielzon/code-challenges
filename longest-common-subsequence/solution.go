package main

func LCS(s1, s2 string) string {
	n, m := len(s1), len(s2)

	// dp[i][j] = length of lcs of s1[0:i] and s2[0:j]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// fill dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// backtrack to build result
	i, j := n, m
	var result []byte

	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			result = append([]byte{s1[i-1]}, result...)
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}