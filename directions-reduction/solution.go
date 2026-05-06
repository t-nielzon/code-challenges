package kata

func dirReduc(arr []string) []string {
	opposite := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}
	stack := make([]string, 0, len(arr))
	for _, d := range arr {
		if n := len(stack); n > 0 && stack[n-1] == opposite[d] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, d)
		}
	}
	return stack
}