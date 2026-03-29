package kata

func Revrot(s string, sz int) string {
	if sz <= 0 || s == "" || sz > len(s) {
		return ""
	}
	result := ""
	for i := 0; i+sz <= len(s); i += sz {
		chunk := s[i : i+sz]
		sum := 0
		for _, c := range chunk {
			d := int(c - '0')
			sum += d * d
		}
		if sum%2 == 0 {
			// reverse
			b := []byte(chunk)
			for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
				b[l], b[r] = b[r], b[l]
			}
			result += string(b)
		} else {
			// rotate left by one
			result += chunk[1:] + chunk[:1]
		}
	}
	return result
}