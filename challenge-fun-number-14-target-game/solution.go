func TargetGame(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	prev2 := 0
	prev1 := max(0, vals[0])

	for i := 1; i < len(vals); i++ {
		current := max(prev1, vals[i]+prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}