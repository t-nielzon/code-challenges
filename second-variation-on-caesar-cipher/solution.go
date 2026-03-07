package kata

func shiftChar(c byte, shift int) byte {
	if c >= 'a' && c <= 'z' {
		return byte((int(c-'a')+shift%26+26)%26) + 'a'
	}
	if c >= 'A' && c <= 'Z' {
		return byte((int(c-'A')+shift%26+26)%26) + 'A'
	}
	return c
}

func Encode(s string, shift int) []string {
	if len(s) == 0 {
		return []string{}
	}

	firstLower := s[0]
	if firstLower >= 'A' && firstLower <= 'Z' {
		firstLower += 32
	}
	prefix := string([]byte{firstLower, shiftChar(firstLower, shift)})

	coded := make([]byte, len(s))
	for i := range s {
		coded[i] = shiftChar(s[i], shift)
	}

	full := prefix + string(coded)
	L := len(full)
	chunk := (L + 4) / 5

	var result []string
	for i := 0; i < 5; i++ {
		start := i * chunk
		if start >= L {
			break
		}
		end := start + chunk
		if end > L {
			end = L
		}
		if end > start {
			result = append(result, full[start:end])
		}
	}
	return result
}

func Decode(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	full := ""
	for _, s := range arr {
		full += s
	}
	if len(full) < 2 {
		return ""
	}

	shift := int(full[1]) - int(full[0])
	if shift < 0 {
		shift += 26
	}

	coded := full[2:]
	decoded := make([]byte, len(coded))
	for i := range coded {
		decoded[i] = shiftChar(coded[i], -shift)
	}
	return string(decoded)
}