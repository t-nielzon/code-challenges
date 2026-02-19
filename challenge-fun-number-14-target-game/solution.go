package kata

func Solve(vals []int) int {
	n := len(vals)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if vals[0] > 0 {
			return vals[0]
		}
		return 0
	}

	prev2 := 0
	prev1 := vals[0]
	if prev1 < 0 {
		prev1 = 0
	}

	for i := 1; i < n; i++ {
		take := prev2 + vals[i]
		skip := prev1

		cur := skip
		if take > cur {
			cur = take
		}
		if cur < 0 {
			cur = 0
		}

		prev2 = prev1
		prev1 = cur
	}

	return prev1
}