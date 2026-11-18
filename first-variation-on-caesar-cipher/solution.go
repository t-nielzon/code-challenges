package kata

import "unicode"

func movingShift(s string, shift int) []string {
	encoded := encode(s, shift)
	return splitMessage(encoded)
}

func demovingShift(s []string, shift int) string {
	joined := ""
	for _, part := range s {
		joined += part
	}
	return decode(joined, shift)
}

func encode(s string, initialShift int) string {
	result := ""
	currentShift := initialShift
	
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) {
				shifted := rune((int(ch-'A') + currentShift) % 26)
				result += string('A' + shifted)
			} else {
				shifted := rune((int(ch-'a') + currentShift) % 26)
				result += string('a' + shifted)
			}
		} else {
			result += string(ch)
		}
		currentShift++
	}
	
	return result
}

func decode(s string, initialShift int) string {
	result := ""
	currentShift := initialShift
	
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) {
				shifted := rune((int(ch-'A') - currentShift%26 + 26) % 26)
				result += string('A' + shifted)
			} else {
				shifted := rune((int(ch-'a') - currentShift%26 + 26) % 26)
				result += string('a' + shifted)
			}
		} else {
			result += string(ch)
		}
		currentShift++
	}
	
	return result
}

func splitMessage(s string) []string {
	n := len(s)
	base := n / 5
	extraForParts1To4 := min(n - 4*base, 4)
	
	parts := make([]string, 5)
	idx := 0
	
	for i := 0; i < 5; i++ {
		var partLen int
		
		if i < extraForParts1To4 {
			partLen = base + 1
		} else if i < 4 {
			partLen = base
		} else {
			// part 5 gets whatever is left
			partLen = n - idx
		}
		
		if partLen > 0 {
			parts[i] = s[idx : idx+partLen]
		} else {
			parts[i] = ""
		}
		idx += partLen
	}
	
	return parts
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}