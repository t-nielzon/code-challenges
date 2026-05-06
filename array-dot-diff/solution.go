package kata

func ArrayDiff(a, b []int) []int {
	remove := make(map[int]struct{}, len(b))
	for _, v := range b {
		remove[v] = struct{}{}
	}
	result := make([]int, 0, len(a))
	for _, v := range a {
		if _, ok := remove[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}