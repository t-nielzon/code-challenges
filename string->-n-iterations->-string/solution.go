package kata

func Iterate(s string, n int) string {
	if n == 0 || len(s) <= 1 {
		return s
	}

	current := s
	seen := make(map[string]int)
	seen[s] = 0

	for i := 1; i <= n; i++ {
		current = applyTransform(current)

		if cycleStart, exists := seen[current]; exists {
			cycleLen := i - cycleStart
			remaining := (n - i) % cycleLen

			for j := 0; j < remaining; j++ {
				current = applyTransform(current)
			}
			return current
		}

		seen[current] = i
	}

	return current
}

func applyTransform(s string) string {
	var even, odd []rune
	for i, ch := range s {
		if i%2 == 0 {
			even = append(even, ch)
		} else {
			odd = append(odd, ch)
		}
	}
	return string(even) + string(odd)
}