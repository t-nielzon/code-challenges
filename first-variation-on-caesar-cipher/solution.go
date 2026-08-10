package main

func movingShift(s string, shift int) []string {
	// Encode the string
	encoded := ""
	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			offset := int(ch - 'a')
			newOffset := (offset + shift + i) % 26
			encoded += string(rune('a' + newOffset))
		} else if ch >= 'A' && ch <= 'Z' {
			offset := int(ch - 'A')
			newOffset := (offset + shift + i) % 26
			encoded += string(rune('A' + newOffset))
		} else {
			encoded += string(ch)
		}
	}

	// Split into 5 parts
	length := len(encoded)
	min_size := length / 5
	total_for_first_4 := (min_size + 1) * 4
	if total_for_first_4 > length {
		total_for_first_4 = length
	}

	base := total_for_first_4 / 4
	rem := total_for_first_4 % 4

	var parts []string
	idx := 0
	for i := 0; i < 4; i++ {
		size := base
		if i < rem {
			size++
		}
		parts = append(parts, encoded[idx:idx+size])
		idx += size
	}
	parts = append(parts, encoded[idx:])

	return parts
}

func demovingShift(s []string, shift int) string {
	// Concatenate all parts
	encoded := ""
	for _, part := range s {
		encoded += part
	}

	// Decode by reversing the shift
	decoded := ""
	for i, ch := range encoded {
		if ch >= 'a' && ch <= 'z' {
			offset := int(ch - 'a')
			newOffset := ((offset - shift - i) % 26 + 26) % 26
			decoded += string(rune('a' + newOffset))
		} else if ch >= 'A' && ch <= 'Z' {
			offset := int(ch - 'A')
			newOffset := ((offset - shift - i) % 26 + 26) % 26
			decoded += string(rune('A' + newOffset))
		} else {
			decoded += string(ch)
		}
	}

	return decoded
}