package main

func Encode(n int) []byte {
	if n == 0 {
		return []byte{0}
	}

	var bytes []byte
	for n > 0 {
		bytes = append([]byte{byte(n & 0x7F)}, bytes...)
		n >>= 7
	}

	for i := 0; i < len(bytes)-1; i++ {
		bytes[i] |= 0x80
	}

	return bytes
}

func Decode(data []byte) int {
	result := 0
	for _, b := range data {
		result = (result << 7) | int(b&0x7F)
		if (b & 0x80) == 0 {
			break
		}
	}
	return result
}