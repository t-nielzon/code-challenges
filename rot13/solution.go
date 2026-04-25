package kata

func Rot13(message string) string {
	out := make([]rune, 0, len(message))
	for _, r := range message {
		switch {
		case r >= 'a' && r <= 'z':
			r = 'a' + (r-'a'+13)%26
		case r >= 'A' && r <= 'Z':
			r = 'A' + (r-'A'+13)%26
		}
		out = append(out, r)
	}
	return string(out)
}