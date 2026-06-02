package kata

func shiftChar(c byte, k int) byte {
	switch {
	case c >= 'a' && c <= 'z':
		return byte('a' + ((int(c-'a')+k)%26+26)%26)
	case c >= 'A' && c <= 'Z':
		return byte('A' + ((int(c-'A')+k)%26+26)%26)
	default:
		return c
	}
}

func movingShift(s string, shift int) []string {
	coded := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		coded[i] = shiftChar(s[i], shift+i)
	}

	n := len(coded)
	size := (n + 4) / 5 // ceil(n/5)

	parts := make([]string, 5)
	for i := 0; i < 5; i++ {
		start := i * size
		if start > n {
			start = n
		}
		end := start + size
		if end > n {
			end = n
		}
		parts[i] = string(coded[start:end])
	}
	return parts
}

func demovingShift(s []string, shift int) string {
	var joined []byte
	for _, part := range s {
		joined = append(joined, []byte(part)...)
	}

	decoded := make([]byte, len(joined))
	for i := 0; i < len(joined); i++ {
		decoded[i] = shiftChar(joined[i], -(shift + i))
	}
	return string(decoded)
}