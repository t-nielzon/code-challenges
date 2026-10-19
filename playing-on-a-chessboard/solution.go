package main

func game(n int) []int {
	numerator := n * n

	if numerator%2 == 0 {
		return []int{numerator / 2}
	}

	return []int{numerator, 2}
}