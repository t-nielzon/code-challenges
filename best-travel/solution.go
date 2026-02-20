package kata

func ChooseBestSum(t, k int, ls []int) int {
	best := -1
	combine(ls, k, 0, 0, t, &best)
	return best
}

func combine(ls []int, k, start, sum, t int, best *int) {
	if k == 0 {
		if sum <= t && sum > *best {
			*best = sum
		}
		return
	}
	for i := start; i <= len(ls)-k; i++ {
		next := sum + ls[i]
		if next <= t {
			combine(ls, k-1, i+1, next, t, best)
		}
	}
}