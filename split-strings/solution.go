package kata

func SplitInParts(s string) []string {
	if len(s)%2 != 0 {
		s += "_"
	}
	result := []string{}
	for i := 0; i < len(s); i += 2 {
		result = append(result, s[i:i+2])
	}
	return result
}