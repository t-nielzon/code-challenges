package main

func movingShift(s string, shift int) []string {
	encoded := ""
	for i, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			shiftAmount := shift + i
			if ch >= 'a' && ch <= 'z' {
				encoded += string(rune((int(ch-'a')+shiftAmount)%26 + int('a')))
			} else {
				encoded += string(rune((int(ch-'A')+shiftAmount)%26 + int('A')))
			}
		} else {
			encoded += string(ch)
		}
	}
	
	n := len(encoded)
	if n == 0 {
		return []string{"", "", "", "", ""}
	}
	
	ceil_base := (n + 4) / 5
	
	sizes := [5]int{}
	sum := 0
	
	for i := 0; i < 4; i++ {
		sizes[i] = ceil_base
		sum += sizes[i]
	}
	
	if sum > n {
		for i := 3; i >= 0 && sum > n; i-- {
			sizes[i]--
			sum--
		}
	}
	
	sizes[4] = n - sum
	
	result := make([]string, 5)
	idx := 0
	for i := 0; i < 5; i++ {
		result[i] = encoded[idx : idx+sizes[i]]
		idx += sizes[i]
	}
	
	return result
}

func demovingShift(s []string, shift int) string {
	encoded := ""
	for _, part := range s {
		encoded += part
	}
	
	decoded := ""
	for i, ch := range encoded {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			shiftAmount := shift + i
			if ch >= 'a' && ch <= 'z' {
				pos := ((int(ch-'a') - shiftAmount) % 26 + 26) % 26
				decoded += string(rune(pos + int('a')))
			} else {
				pos := ((int(ch-'A') - shiftAmount) % 26 + 26) % 26
				decoded += string(rune(pos + int('A')))
			}
		} else {
			decoded += string(ch)
		}
	}
	
	return decoded
}