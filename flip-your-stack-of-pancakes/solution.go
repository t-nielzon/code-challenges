package kata

func Solve(stack []int) []int {
	s := make([]int, len(stack))
	copy(s, stack)
	result := []int{}

	for size := len(s) - 1; size > 0; size-- {
		maxIdx := 0
		for i := 1; i <= size; i++ {
			if s[i] > s[maxIdx] {
				maxIdx = i
			}
		}

		if maxIdx == size {
			continue
		}

		if maxIdx > 0 {
			reverse(s, 0, maxIdx)
			result = append(result, maxIdx)
		}

		reverse(s, 0, size)
		result = append(result, size)
	}

	return result
}

func reverse(s []int, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}