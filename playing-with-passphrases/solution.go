package kata

func PlayPass(s string, n int) string {
	runes := []rune(s)
	out := make([]rune, len(runes))
	for i, r := range runes {
		var c rune
		switch {
		case r >= 'A' && r <= 'Z':
			c = 'A' + (r-'A'+rune(n))%26
		case r >= 'a' && r <= 'z':
			c = 'a' + (r-'a'+rune(n))%26
		case r >= '0' && r <= '9':
			c = '9' - (r - '0')
		default:
			c = r
		}
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			if i%2 == 0 {
				if c >= 'a' && c <= 'z' {
					c = c - 'a' + 'A'
				}
			} else {
				if c >= 'A' && c <= 'Z' {
					c = c - 'A' + 'a'
				}
			}
		}
		out[len(runes)-1-i] = c
	}
	return string(out)
}