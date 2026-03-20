package kata

func Rot13(msg string) string {
	result := make([]byte, len(msg))
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		switch {
		case c >= 'a' && c <= 'z':
			result[i] = 'a' + (c-'a'+13)%26
		case c >= 'A' && c <= 'Z':
			result[i] = 'A' + (c-'A'+13)%26
		default:
			result[i] = c
		}
	}
	return string(result)
}