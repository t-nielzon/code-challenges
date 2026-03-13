package kata

func LongestComb(arr []int, command string) [][]int {
	n := len(arr)
	if n < 3 {
		return [][]int{}
	}

	increasing := command == "< <" || command == "<<"

	cmp := func(a, b int) bool {
		if increasing {
			return a < b
		}
		return a > b
	}

	// dp[i] = length of longest valid subsequence ending at index i
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if cmp(arr[j], arr[i]) && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	maxLen := 0
	for _, v := range dp {
		if v > maxLen {
			maxLen = v
		}
	}

	if maxLen < 3 {
		return [][]int{}
	}

	// backtrack to find all subsequences of maxLen
	var results [][]int
	var backtrack func(idx, remaining int, path []int)
	backtrack = func(idx, remaining int, path []int) {
		if remaining == 0 {
			seq := make([]int, len(path))
			copy(seq, path)
			results = append(results, seq)
			return
		}
		for i := idx; i < n; i++ {
			if dp[i] != remaining {
				continue
			}
			if len(path) == 0 || cmp(path[len(path)-1], arr[i]) {
				// check if we can still complete from here
				backtrack(i+1, remaining-1, append(path, arr[i]))
			}
		}
	}

	backtrack(0, maxLen, nil)
	return results
}