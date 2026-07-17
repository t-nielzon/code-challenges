package main

func Score(dice []int) int {
	freq := make(map[int]int)
	for _, d := range dice {
		freq[d]++
	}

	score := 0
	for value := 1; value <= 6; value++ {
		count := freq[value]

		if count >= 3 {
			if value == 1 {
				score += 1000
			} else {
				score += value * 100
			}
			count -= 3
		}

		if count > 0 {
			if value == 1 {
				score += count * 100
			} else if value == 5 {
				score += count * 50
			}
		}
	}

	return score
}