package kata

func ArrayDiff(a, b []int) []int {
	exclude := make(map[int]bool, len(b))
	for _, v := range b {
		exclude[v] = true
	}
	result := make([]int, 0, len(a))
	for _, v := range a {
		if !exclude[v] {
			result = append(result, v)
		}
	}
	return result
}