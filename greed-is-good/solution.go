package kata

func Score(dice [5]int) int {
	counts := [7]int{}
	for _, d := range dice {
		counts[d]++
	}

	score := 0

	// triplet scores
	tripletPoints := [7]int{0, 1000, 200, 300, 400, 500, 600}
	for face := 1; face <= 6; face++ {
		if counts[face] >= 3 {
			score += tripletPoints[face]
			counts[face] -= 3
		}
	}

	// single scoring dice
	score += counts[1] * 100
	score += counts[5] * 50

	return score
}