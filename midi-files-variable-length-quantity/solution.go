package kata

func Encode(n int) []byte {
	if n <= 127 {
		return []byte{byte(n)}
	}

	var result []byte

	for n > 0 {
		result = append(result, byte(n&0x7F))
		n >>= 7
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	for i := 0; i < len(result)-1; i++ {
		result[i] |= 0x80
	}

	return result
}

func Decode(data []byte) int {
	result := 0

	for _, b := range data {
		result = (result << 7) | int(b&0x7F)

		if b&0x80 == 0 {
			break
		}
	}

	return result
}