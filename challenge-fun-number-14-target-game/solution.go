package kata

func MaxPoints(vals []int) int {
	prev, curr := 0, 0
	for _, v := range vals {
		next := curr
		if prev+v > next {
			next = prev + v
		}
		prev, curr = curr, next
	}
	return curr
}