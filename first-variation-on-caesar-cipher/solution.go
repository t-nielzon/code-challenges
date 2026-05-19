package kata

import "strings"

func shiftChar(c byte, shift int) byte {
	switch {
	case c >= 'a' && c <= 'z':
		return byte((int(c-'a')+shift%26+26)%26) + 'a'
	case c >= 'A' && c <= 'Z':
		return byte((int(c-'A')+shift%26+26)%26) + 'A'
	default:
		return c
	}
}

func MovingShift(s string, shift int) [5]string {
	var coded strings.Builder
	for i := 0; i < len(s); i++ {
		coded.WriteByte(shiftChar(s[i], shift+i))
	}
	encoded := coded.String()
	n := len(encoded)
	base := n / 5
	rem := n % 5
	var result [5]string
	idx := 0
	for i := 0; i < 5; i++ {
		size := base
		if i < rem {
			size++
		}
		result[i] = encoded[idx : idx+size]
		idx += size
	}
	return result
}

func DemovingShift(s [5]string, shift int) string {
	joined := strings.Join(s[:], "")
	var decoded strings.Builder
	for i := 0; i < len(joined); i++ {
		decoded.WriteByte(shiftChar(joined[i], -(shift + i)))
	}
	return decoded.String()
}