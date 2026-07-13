package kata

func SplitStrings(str string) []string {
	var result []string
	runes := []rune(str)
	
	for i := 0; i < len(runes); i += 2 {
		if i+1 < len(runes) {
			result = append(result, string(runes[i:i+2]))
		} else {
			result = append(result, string(runes[i])+"_")
		}
	}
	
	return result
}