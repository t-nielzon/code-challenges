package main

func Score(dice []int) int {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}

	score := 0
	tripletScores := map[int]int{
		1: 1000,
		2: 200,
		3: 300,
		4: 400,
		5: 500,
		6: 600,
	}

	for val := 1; val <= 6; val++ {
		count := counts[val]
		if count >= 3 {
			score += tripletScores[val]
			count -= 3
		}
		if val == 1 {
			score += count * 100
		} else if val == 5 {
			score += count * 50
		}
	}

	return score
}