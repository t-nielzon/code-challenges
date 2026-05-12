package kata

func TargetGame(vals []int) int {
	prev, curr := 0, 0
	for _, v := range vals {
		prev, curr = curr, max(curr, prev+v)
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}