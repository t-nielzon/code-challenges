package kata

func encode(n int) []byte {
	// collect 7-bit groups least-significant first, then reverse
	groups := []byte{byte(n & 0x7F)}
	for n >>= 7; n > 0; n >>= 7 {
		groups = append(groups, byte(n&0x7F))
	}

	out := make([]byte, len(groups))
	for i, j := 0, len(groups)-1; j >= 0; i, j = i+1, j-1 {
		// set continuation bit on every byte except the last
		if i < len(groups)-1 {
			out[i] = groups[j] | 0x80
		} else {
			out[i] = groups[j]
		}
	}
	return out
}

func decode(data []byte) int {
	result := 0
	for _, b := range data {
		result = (result << 7) | int(b&0x7F)
		// highest bit clear marks the final byte of the first VLQ
		if b&0x80 == 0 {
			break
		}
	}
	return result
}