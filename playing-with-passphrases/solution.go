package kata

func PlayPass(s string, n int) string {
	runes := []rune(s)
	out := make([]rune, len(runes))
	for i, r := range runes {
		switch {
		case r >= 'A' && r <= 'Z':
			shifted := rune((int(r-'A')+n)%26) + 'A'
			if i%2 == 1 {
				shifted = shifted - 'A' + 'a'
			}
			out[i] = shifted
		case r >= 'a' && r <= 'z':
			shifted := rune((int(r-'a')+n)%26) + 'a'
			if i%2 == 0 {
				shifted = shifted - 'a' + 'A'
			}
			out[i] = shifted
		case r >= '0' && r <= '9':
			out[i] = '9' - (r - '0')
		default:
			out[i] = r
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}