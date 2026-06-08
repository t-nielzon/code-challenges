package kata

func Score(dice []int) int {
	counts := make([]int, 7)
	for _, d := range dice {
		counts[d]++
	}

	score := 0
	for face := 1; face <= 6; face++ {
		n := counts[face]

		if n >= 3 {
			if face == 1 {
				score += 1000
			} else {
				score += face * 100
			}
			n -= 3
		}

		if face == 1 {
			score += n * 100
		} else if face == 5 {
			score += n * 50
		}
	}

	return score
}