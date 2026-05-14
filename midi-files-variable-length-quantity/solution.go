package kata

func Encode(n int) []byte {
	if n == 0 {
		return []byte{0x00}
	}
	var buf []byte
	for n > 0 {
		buf = append([]byte{byte(n & 0x7F)}, buf...)
		n >>= 7
	}
	for i := 0; i < len(buf)-1; i++ {
		buf[i] |= 0x80
	}
	return buf
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