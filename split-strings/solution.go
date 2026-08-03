package kata

func Solution(s string) []string {
	if len(s)%2 == 1 {
		s += "_"
	}
	
	var result []string
	for i := 0; i < len(s); i += 2 {
		result = append(result, s[i:i+2])
	}
	
	return result
}