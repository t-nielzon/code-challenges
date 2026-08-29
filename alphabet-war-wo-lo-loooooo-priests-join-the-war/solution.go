// solution.go
package main

func AlphabetWar(fight string) string {
	s := []rune(fight)
	n := len(s)

	leftLetters := map[rune]int{'w': 4, 'p': 3, 'b': 2, 's': 1, 't': 0}
	rightLetters := map[rune]int{'m': 4, 'q': 3, 'd': 2, 'z': 1, 'j': 0}

	isLeftSide := func(r rune) bool {
		_, ok := leftLetters[r]
		return ok
	}

	isRightSide := func(r rune) bool {
		_, ok := rightLetters[r]
		return ok
	}

	hasBothPriests := func(i int) bool {
		hasT := (i > 0 && s[i-1] == 't') || (i < n-1 && s[i+1] == 't')
		hasJ := (i > 0 && s[i-1] == 'j') || (i < n-1 && s[i+1] == 'j')
		return hasT && hasJ
	}

	for i := 0; i < n; i++ {
		if s[i] == 't' {
			if i > 0 && isRightSide(s[i-1]) && s[i-1] != 'j' && !hasBothPriests(i-1) {
				power := rightLetters[s[i-1]]
				for k, v := range leftLetters {
					if v == power && k != 't' {
						s[i-1] = k
						break
					}
				}
			}
			if i < n-1 && isRightSide(s[i+1]) && s[i+1] != 'j' && !hasBothPriests(i+1) {
				power := rightLetters[s[i+1]]
				for k, v := range leftLetters {
					if v == power && k != 't' {
						s[i+1] = k
						break
					}
				}
			}
		} else if s[i] == 'j' {
			if i > 0 && isLeftSide(s[i-1]) && s[i-1] != 't' && !hasBothPriests(i-1) {
				power := leftLetters[s[i-1]]
				for k, v := range rightLetters {
					if v == power && k != 'j' {
						s[i-1] = k
						break
					}
				}
			}
			if i < n-1 && isLeftSide(s[i+1]) && s[i+1] != 't' && !hasBothPriests(i+1) {
				power := leftLetters[s[i+1]]
				for k, v := range rightLetters {
					if v == power && k != 'j' {
						s[i+1] = k
						break
					}
				}
			}
		}
	}

	leftPower := 0
	rightPower := 0

	for _, r := range s {
		if power, ok := leftLetters[r]; ok {
			leftPower += power
		} else if power, ok := rightLetters[r]; ok {
			rightPower += power
		}
	}

	if leftPower > rightPower {
		return "Left side wins!"
	} else if rightPower > leftPower {
		return "Right side wins!"
	} else {
		return "Let's fight again!"
	}
}