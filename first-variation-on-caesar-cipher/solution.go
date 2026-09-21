package main

func movingShift(s string, shift int) []string {
	encoded := ""
	shift_counter := shift
	
	for _, char := range s {
		if 'a' <= char && char <= 'z' {
			offset := (int(char-'a') + shift_counter) % 26
			encoded += string('a' + rune(offset))
		} else if 'A' <= char && char <= 'Z' {
			offset := (int(char-'A') + shift_counter) % 26
			encoded += string('A' + rune(offset))
		} else {
			encoded += string(char)
		}
		shift_counter++
	}
	
	return splitMessage(encoded)
}

func demovingShift(s []string, shift int) string {
	encoded := ""
	for _, part := range s {
		encoded += part
	}
	
	decoded := ""
	shift_counter := shift
	
	for _, char := range encoded {
		if 'a' <= char && char <= 'z' {
			offset := ((int(char-'a') - shift_counter%26) % 26 + 26) % 26
			decoded += string('a' + rune(offset))
		} else if 'A' <= char && char <= 'Z' {
			offset := ((int(char-'A') - shift_counter%26) % 26 + 26) % 26
			decoded += string('A' + rune(offset))
		} else {
			decoded += string(char)
		}
		shift_counter++
	}
	
	return decoded
}

func splitMessage(msg string) []string {
	length := len(msg)
	ceil_size := (length + 4) / 5
	
	sizes := []int{ceil_size, ceil_size, ceil_size, ceil_size, ceil_size}
	
	total := 5 * ceil_size
	for total > length {
		for i := 4; i >= 0; i-- {
			if sizes[i] > 0 {
				sizes[i]--
				total--
				break
			}
		}
	}
	
	result := make([]string, 5)
	pos := 0
	for i := 0; i < 5; i++ {
		result[i] = msg[pos : pos+sizes[i]]
		pos += sizes[i]
	}
	
	return result
}