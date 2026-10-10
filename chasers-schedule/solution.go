package main

func MaxDistance(s, t int) int {
	// dp[i][speed][state]
	// state 0 = can sprint, state 1 = must recover
	dp := make([][][]int, t+1)
	for i := 0; i <= t; i++ {
		dp[i] = make([][]int, s+1)
		for j := 0; j <= s; j++ {
			dp[i][j] = []int{-1, -1}
		}
	}

	dp[0][s][0] = 0

	for i := 0; i < t; i++ {
		for speed := 0; speed <= s; speed++ {
			// state 0: can sprint
			if dp[i][speed][0] >= 0 {
				// option 1: normal run
				newDist := dp[i][speed][0] + speed
				if dp[i+1][speed][0] < newDist {
					dp[i+1][speed][0] = newDist
				}

				// option 2: sprint
				if speed > 0 {
					newDist = dp[i][speed][0] + speed*2
					if dp[i+1][speed-1][1] < newDist {
						dp[i+1][speed-1][1] = newDist
					}
				}
			}

			// state 1: must recover
			if dp[i][speed][1] >= 0 {
				newDist := dp[i][speed][1] + speed
				if dp[i+1][speed][0] < newDist {
					dp[i+1][speed][0] = newDist
				}
			}
		}
	}

	result := 0
	for speed := 0; speed <= s; speed++ {
		if dp[t][speed][0] >= 0 {
			result = max(result, dp[t][speed][0])
		}
		if dp[t][speed][1] >= 0 {
			result = max(result, dp[t][speed][1])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}