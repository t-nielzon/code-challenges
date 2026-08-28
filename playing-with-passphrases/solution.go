package main

func PlayPassphrase(s string, shift int) string {
	// Step 1: Shift letters circularly
	shifted := make([]rune, 0)
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			idx := int(r-'A') + shift
			idx = ((idx % 26) + 26) % 26
			shifted = append(shifted, rune('A'+idx))
		} else if r >= '0' && r <= '9' {
			shifted = append(shifted, r)
		} else {
			shifted = append(shifted, r)
		}
	}
	
	// Step 2: Replace digits with complement to 9
	for i, r := range shifted {
		if r >= '0' && r <= '9' {
			digit := int(r - '0')
			shifted[i] = rune('0' + (9-digit))
		}
	}
	
	// Step 4: Downcase odd positions, upcase even positions
	for i, r := range shifted {
		if r >= 'A' && r <= 'Z' {
			if i%2 == 1 {
				shifted[i] = r + 32
			}
		}
	}
	
	// Step 5: Reverse
	for i, j := 0, len(shifted)-1; i < j; i, j = i+1, j-1 {
		shifted[i], shifted[j] = shifted[j], shifted[i]
	}
	
	return string(shifted)
}