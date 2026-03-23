package kata

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += "_"
	}
	result := make([]string, 0, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return result
}