package kata

func ArrayDiff(a, b []int) []int {
	// build a set from b for O(1) lookups
	exclude := make(map[int]bool)
	for _, v := range b {
		exclude[v] = true
	}

	result := make([]int, 0)
	for _, v := range a {
		if !exclude[v] {
			result = append(result, v)
		}
	}

	return result
}