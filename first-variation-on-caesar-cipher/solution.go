package kata

func shiftRune(c byte, amount int) byte {
	amount = ((amount % 26) + 26) % 26
	switch {
	case c >= 'a' && c <= 'z':
		return 'a' + (c-'a'+byte(amount))%26
	case c >= 'A' && c <= 'Z':
		return 'A' + (c-'A'+byte(amount))%26
	default:
		return c
	}
}

func movingShift(s string, shift int) []string {
	coded := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		coded[i] = shiftRune(s[i], shift+i)
	}

	n := len(coded)
	// ceil(n/5) gives the size of the leading, fully-packed parts
	chunk := (n + 4) / 5

	parts := make([]string, 5)
	pos := 0
	for p := 0; p < 5; p++ {
		end := pos + chunk
		if end > n {
			end = n
		}
		parts[p] = string(coded[pos:end])
		pos = end
	}
	return parts
}

func demovingShift(s []string, shift int) string {
	joined := ""
	for _, part := range s {
		joined += part
	}

	decoded := make([]byte, len(joined))
	for i := 0; i < len(joined); i++ {
		decoded[i] = shiftRune(joined[i], -(shift + i))
	}
	return string(decoded)
}