package kata

func CleanString(s string) string {
	result := []rune{}
	for _, c := range s {
		if c == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}