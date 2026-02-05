package kata

func Solution(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	// Pad with underscore if odd length
	if len(str)%2 != 0 {
		str += "_"
	}

	result := make([]string, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		result[i/2] = str[i : i+2]
	}

	return result
}