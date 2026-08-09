package main

func ROT13(input string) string {
	result := []rune(input)
	for i, char := range result {
		switch {
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		}
	}
	return string(result)
}