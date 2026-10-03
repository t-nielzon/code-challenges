package main

func Fight(fight string) string {
	runes := []rune(fight)
	killed := make(map[int]bool)

	// mark positions killed by bombs
	for i, r := range runes {
		if r == '*' {
			if i > 0 {
				killed[i-1] = true
			}
			if i < len(runes)-1 {
				killed[i+1] = true
			}
		}
	}

	leftPower := map[rune]int{'w': 4, 'p': 3, 'b': 2, 's': 1}
	rightPower := map[rune]int{'m': 4, 'q': 3, 'd': 2, 'z': 1}

	leftScore := 0
	rightScore := 0

	for i, r := range runes {
		if !killed[i] && r != '*' {
			if p, ok := leftPower[r]; ok {
				leftScore += p
			} else if p, ok := rightPower[r]; ok {
				rightScore += p
			}
		}
	}

	if leftScore > rightScore {
		return "Left side wins!"
	} else if rightScore > leftScore {
		return "Right side wins!"
	} else {
		return "Let's fight again!"
	}
}